package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
)

type CommentCount struct {
    Id         uint64
    PraiseNum  uint32
    ReplyNum   uint32
    DislikeNum uint32
}

type CommentCountRepo interface {
    CommentInfo(ctx context.Context, commentId uint64) (*CommentCount, error)
}

type CommentCountUseCase struct {
    repo CommentCountRepo
    log  *log.Helper
}

func (uc *CommentCountUseCase) CommentInfo(ctx context.Context, commentId uint64) (*CommentCount, error) {
    return uc.repo.CommentInfo(ctx, commentId)
}
