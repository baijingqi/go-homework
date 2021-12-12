package service

import (
    v1 "comment/api/comment"
    "comment/internal/biz"
    "context"
    "github.com/go-kratos/kratos/v2/log"
)

// CommentService is a Comment service.
type CommentService struct {
    v1.UnimplementedCommentServer

    cc  *biz.CommentUseCase
    log *log.Helper
}

func NewCommentService(oc *biz.CommentUseCase, logger log.Logger) *CommentService {
    return &CommentService{
        cc:  oc,
        log: log.NewHelper(log.With(logger, "module", "service/comment"))}
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
    if err != nil {
        return &v1.AddCommentReply{
            Id: 0,
        }, err
    } else {
        return &v1.AddCommentReply{
            Id: 1,
        }, err
    }

}
