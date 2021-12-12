package data

import "C"
import (
    "comment/internal/biz"
    "context"
    "github.com/go-kratos/kratos/v2/log"
)

type CommentCount struct {
    Id         uint64
    PraiseNum  uint32
    ReplyNum   uint32
    DislikeNum uint32
}

type CommentCountRepo struct {
    data *Data
    log  *log.Helper
}

// NewCommentCountRepo .
func NewCommentCountRepo(data *Data, logger log.Logger) biz.CommentCountRepo {
    return &CommentCountRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}
func (uc *CommentCountRepo) CommentInfo(ctx context.Context, commentId uint64) (*biz.CommentCount, error) {
    arr, err := uc.data.DB.CommentCount.Get(ctx, commentId)
    if err != nil {
        return  &biz.CommentCount{},err
    }
    return &biz.CommentCount{
        Id:arr.ID,
        PraiseNum:arr.PraiseNum,
        DislikeNum:arr.DislikeNum,
        ReplyNum:arr.ReplyNum,
    }, err
}
