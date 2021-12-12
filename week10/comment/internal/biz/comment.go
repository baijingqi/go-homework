package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "time"
)

type Comment struct {
    Id              uint64
    Uid             uint64
    RelationId      uint64
    RelationType    uint32
    ParentId        uint64
    Content         string
    BelongCommentId uint64
    PraiseNum       uint32
    ReplyNum        uint32
    CreatedAt       time.Time
}

type CommentRepo interface {
    CreateComment(context.Context, *Comment) error
    DelComment(context.Context, *Comment) (bool, error)
    CommentList(ctx context.Context, commentId uint64, relationId uint64, relationType uint32, uid uint64, page uint32, size uint) ([]*Comment, error)
}

type CommentUseCase struct {
    repo CommentRepo
    log  *log.Helper
}

func NewCommentUseCase(repo CommentRepo, logger log.Logger) *CommentUseCase {
    return &CommentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUseCase) Create(ctx context.Context, g *Comment) error {
    return uc.repo.CreateComment(ctx, g)
}

func (uc *CommentUseCase) Del(ctx context.Context, g *Comment) (bool, error) {
    return uc.repo.DelComment(ctx, g)
}

func (uc *CommentUseCase) CommentList(ctx context.Context, commentId uint64, relationId uint64, relationType uint32, uid uint64, page uint32, size uint) ([]*Comment, error) {
    arr, err := uc.repo.CommentList(ctx, commentId, relationId, relationType, uid, page, size)
    if err != nil {
        uc.log.Error(log.LevelFatal, err)
        return arr, err
    }
    countCase := &CommentCountUseCase{}
    for key, val := range arr {
        info, err := countCase.CommentInfo(ctx, val.Id)
        if err == nil {
            arr[key].PraiseNum = info.PraiseNum
            arr[key].ReplyNum = info.ReplyNum
        }
    }
    return arr, err
}
