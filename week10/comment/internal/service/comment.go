package service

import (
    v1 "comment/api/comment"
    pb "comment/api/user"
    "comment/internal/biz"
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "google.golang.org/grpc"
)


// CommentService is a Comment service.
type CommentService struct {
    v1.UnimplementedCommentServer

    cc        *biz.CommentUseCase
    countCase *biz.CommentCountUseCase
    log       *log.Helper
}

func NewCommentService(oc *biz.CommentUseCase, countCase *biz.CommentCountUseCase, logger log.Logger) *CommentService {
    return &CommentService{
        cc:        oc,
        countCase: countCase,
        log:       log.NewHelper(log.With(logger, "module", "service/comment"))}
}

func (c *CommentService) CommentList(ctx context.Context, req *v1.CommentListRequest) (*v1.CommentListReply, error) {
    if req.RelationId == 0 || req.RelationType == 0 || req.Page == 0 {
        return &v1.CommentListReply{
            Code: 417,
            Msg:  "请求参数错误",
        }, nil
    }

    arr, err := c.cc.CommentList(ctx, c.countCase, req.CommentId, req.RelationId, req.RelationType, req.Uid, req.Page, 20)
    if err != nil {
        return &v1.CommentListReply{
            Code: 500,
            Msg:  "获取列表失败",
        }, nil
    }

    list := convertCommentList(arr)

    for key, val := range list {
        if val.ParentId == 0 {
            replyList, err := c.cc.CommentList(ctx, c.countCase, val.Id, 0, 0, 0, 1, 3)
            if err != nil {
                continue
            }
            list[key].ReplyList = convertCommentList(replyList)
        }
    }
    return &v1.CommentListReply{
        Code: 200,
        Msg:  "获取列表成功",
        List: list,
    }, nil
}

func convertCommentList(arr []*biz.Comment) []*v1.CommentListStruct {
    var userIds []uint64
    for _, val := range arr {
        userIds = append(userIds, val.Uid)
    }
    users := getUsers(userIds)

    var list []*v1.CommentListStruct
    for _, val := range arr {
        var replyList []*v1.CommentListStruct
        currentUser := &pb.UserStruct{}
        if IssetUint64(val.Uid, users) {
            currentUser = users[val.Uid]
        }
        list = append(list, &v1.CommentListStruct{
            Id:        val.Id,
            Content:   val.Content,
            ParentId:  val.ParentId,
            CreatedAt: val.CreatedAt.Format("2006-01-02 15:04:05"),
            PraiseNum: uint64(val.PraiseNum),
            ReplyNum:  uint64(val.ReplyNum),
            ReplyList: replyList,
            User: &v1.UserStruct{
                Uid:      currentUser.Uid,
                Level:    currentUser.Level,
                Nickname: currentUser.Nickname,
                Avatar:   currentUser.Avatar,
            },
        })
    }
    return list
}

func getUsers(userIds []uint64) map[uint64]*pb.UserStruct {
    conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
    if err != nil {
        return nil
    }
    defer conn.Close()
    c := pb.NewUserClient(conn)
    req := &pb.UserListRequest{
        UserIds: userIds,
    }
    r, err := c.UserList(context.Background(), req)
    if err != nil {
        return nil
    }
    res := make(map[uint64]*pb.UserStruct, len(userIds))
    for _, val := range r.List {
        res[val.Uid] = val
    }
    return res
}

func (c *CommentService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (*v1.AddCommentReply, error) {
    comment := &biz.Comment{
        Uid:          req.GetUid(),
        Content:      req.GetContent(),
        ParentId:     req.GetParentId(),
        RelationId:   req.GetRelationId(),
        RelationType: req.GetRelationType(),
    }
    err := c.cc.Create(ctx, comment)
    ok := true
    if err != nil {
        ok = false
    }
    return &v1.AddCommentReply{
        Ok: ok,
    }, err
}

func IssetUint64(num uint64, arr map[uint64]*pb.UserStruct) bool {
    if _, ok := arr[num]; ok {
        return true
    }
    return false
}
