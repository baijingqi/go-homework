package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
)

type CommentCount struct {
    Id         uint64 `json:"Id,omitempty"`
    PraiseNum  uint32 `json:"PraiseNum"`
    ReplyNum   uint32 `json:"ReplyNum"`
    DislikeNum uint32 `json:"DislikeNum"`
}

type CommentCountRepo interface {
    CommentInfo(ctx context.Context, commentId uint64) (*CommentCount, error)
    BatchCommentInfo(ctx context.Context, commentIds []uint64) (map[uint64]*CommentCount, error)
}

type CommentCountUseCase struct {
    repo CommentCountRepo
    log  *log.Helper
}

func NewCommentCountUseCase(repo CommentCountRepo, logger log.Logger) *CommentCountUseCase {
    return &CommentCountUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentCountUseCase) CommentInfo(ctx context.Context, commentId uint64) (*CommentCount, error) {
    res, err := uc.repo.CommentInfo(ctx, commentId)
    return res, err
}
func (uc *CommentCountUseCase) BatchCommentInfo(ctx context.Context, commentIds []uint64) (map[uint64]*CommentCount, error) {
    res, err := uc.repo.BatchCommentInfo(ctx, commentIds)
    return res, err
}
