package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
)

type Comment struct {
    Id           uint64
    Uid          uint64
    RelationId   uint64
    RelationType uint32
    ParentId     uint64
    Content      string
}

type CommentRepo interface {
    CreateComment(context.Context, *Comment) error
    DelComment(context.Context, *Comment) (bool, error)
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
