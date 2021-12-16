package data

import "C"
import (
    "comment/ent"
    "comment/internal/biz"
    "context"
    "errors"
    "github.com/fatih/structs"
    "github.com/go-kratos/kratos/v2/log"
    "math/rand"
    "strconv"
    "time"
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
    cacheKey := "commentCountInfo:" + strconv.FormatUint(commentId, 10)
    cmd := uc.data.Redis.HGetAll(cacheKey)
    result, err := cmd.Result()
    if err != nil {
        return nil, err
    }
    if len(result) <= 0 {
        arr, err := uc.data.DB.CommentCount.Get(ctx, commentId)
        if err != nil {
            arr = &ent.CommentCount{}
        }
        commentMap := structs.Map(arr)
        commentMap["ID"] = commentId

        notFound := false
        var notFoundError *ent.NotFoundError
        if err != nil {
            if errors.As(err, &notFoundError) {
                notFound = true
            } else {
                return &biz.CommentCount{}, err
            }
        }
        uc.data.Redis.HMSet(cacheKey, commentMap)
        rand.Seed(time.Now().UnixNano())
        if notFound {
            uc.data.Redis.Expire(cacheKey, 100*time.Second)
        } else {
            uc.data.Redis.Expire(cacheKey, time.Duration(86400+rand.Intn(7200))*time.Second)
        }
        return &biz.CommentCount{
            Id:         arr.ID,
            PraiseNum:  arr.PraiseNum,
            DislikeNum: arr.DislikeNum,
            ReplyNum:   arr.ReplyNum,
        }, err
    } else {
        info := biz.CommentCount{}
        num, err := strconv.ParseInt(result["PraiseNum"], 10, 64)
        info.PraiseNum = uint32(num)

        dislikeNum, err := strconv.ParseInt(result["DislikeNum"], 10, 64)
        info.DislikeNum = uint32(dislikeNum)

        replyNum, err := strconv.ParseInt(result["ReplyNum"], 10, 64)
        info.ReplyNum = uint32(replyNum)
        return &info, err
    }
}
