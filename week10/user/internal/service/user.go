package service

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "strconv"
    "user/internal/biz"

    pb "user/api/user"
)

type UserService struct {
    pb.UnimplementedUserServer
    uc  *biz.UserUsecase
    log *log.Helper
}

func NewUserService(oc *biz.UserUsecase, logger log.Logger) *UserService {
    return &UserService{
        uc:  oc,
        log: log.NewHelper(log.With(logger, "module", "service/user"))}
}
func (s *UserService) UserList(ctx context.Context, req *pb.UserListRequest) (*pb.UserListReply, error) {
    arr := make([]*pb.UserStruct, 0)

    for _, userId := range req.UserIds {
        arr = append(arr, &pb.UserStruct{
            Uid:      userId,
            Avatar:   "https://www.bilibili.com/" + strconv.FormatUint(userId, 10) + ".jpg",
            Level:    1,
            Nickname: "道友" + strconv.FormatUint(userId, 10),
        })
    }
    return &pb.UserListReply{
        Code: 200,
        Msg:  "success",
        List: arr,
    }, nil
}
