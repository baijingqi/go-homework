package data

import "C"
import (
    entComment "comment/ent/comment"
    "comment/internal/biz"
    "context"
    "encoding/json"
    "github.com/Shopify/sarama"
    "github.com/go-kratos/kratos/v2/log"
)

type Comment struct {
    Uid          uint64 `json:"uid"`
    ParentId     uint64 `json:"parentId"`
    RelationId   uint64 `json:"relationId"`
    RelationType uint32 `json:"RelationType"`
    Content      string `json:"content"`
}

type CommentRepo struct {
    data *Data
    log  *log.Helper
}

// NewCommentRepo .
func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
    return &CommentRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (r *CommentRepo) CreateComment(ctx context.Context, g *biz.Comment) error {
    r.data.db.Debug()
    msg := &sarama.ProducerMessage{}
    msg.Topic = "comment"

    comment := &Comment{
        Uid:          g.Uid,
        ParentId:     g.ParentId,
        RelationId:   g.RelationId,
        RelationType: g.RelationType,
        Content:      g.Content,
    }
    jsonStr, _ := json.Marshal(comment)
    msg.Value = sarama.StringEncoder(jsonStr)

    _, _, err := r.data.Kafka.SendMessage(msg)
    if err != nil {
        r.log.Error("添加评论失败 jsonStr:", jsonStr)
    }
    return err
}

func (r *CommentRepo) UpdateComment(ctx context.Context, g *biz.Comment) error {
    return nil
}
func (r *CommentRepo) DelComment(ctx context.Context, g *biz.Comment) (bool, error) {
    return true, nil
}

func (r *CommentRepo) CommentList(ctx context.Context, commentId uint64, relationId uint64, relationType uint32, uid uint64, page uint32, size uint) ([]*biz.Comment, error) {
    var arr []*biz.Comment
    client := r.data.DB.Comment
    query := client.Query().Where(entComment.RelationIDEQ(relationId)).Where(entComment.RelationTypeEQ(relationType))
    if commentId != 0 {
        query = query.Where(entComment.BelongCommentIDEQ(commentId))
    }
    if uid != 0 {
        query = query.Where(entComment.UIDEQ(uid))
    }
    limit := (page - 1) * uint32(size)
    res, err := query.Limit(int(limit)).Offset(int(size)).All(ctx)
    if err != nil {
        r.log.Error(err)
        return arr, err
    }
    for _, val := range res {
        arr = append(arr, &biz.Comment{
            Id:              val.ID,
            Uid:             val.UID,
            RelationId:      val.RelationID,
            RelationType:    val.RelationType,
            BelongCommentId: val.BelongCommentID,
            CreatedAt:       val.CreatedAt,
        })
    }
    return arr, err
}
