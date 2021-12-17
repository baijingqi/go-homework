package data

import "C"
import (
    "comment/ent"
    "comment/internal/biz"
    "comment/pkg"
    "context"
    "errors"
    "github.com/fatih/structs"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-redis/redis"
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
        info.PraiseNum = pkg.ToUint32(result["PraiseNum"])
        info.DislikeNum = pkg.ToUint32(result["DislikeNum"])
        info.ReplyNum = pkg.ToUint32(result["ReplyNum"])
        return &info, err
    }
}

func (uc *CommentCountRepo) BatchCommentInfo(ctx context.Context, commentIds []uint64) (map[uint64]*biz.CommentCount, error) {
    pipeline := uc.data.Redis.Pipeline()
    for _, commentId := range commentIds {
        pipeline.HGetAll("commentCountInfo:" + strconv.FormatUint(commentId, 10))
    }

    result := make(map[uint64]*biz.CommentCount)
    cmd, _ := pipeline.Exec()

    for _, val := range cmd {
        mapCmd := val.(*redis.StringStringMapCmd)
        resMap := mapCmd.Val()
        id := pkg.ToUint64(resMap["ID"])
        result[id] = &biz.CommentCount{
            Id:         id,
            PraiseNum:  pkg.ToUint32(resMap["PraiseNum"]),
            ReplyNum:   pkg.ToUint32(resMap["ReplyNum"]),
            DislikeNum: pkg.ToUint32(resMap["DislikeNum"]),
        }
    }

    reQueryIds := make([]uint64, 0)
    for _, commentId := range commentIds {
        if _, ok := result[commentId]; !ok {
            reQueryIds = append(reQueryIds, commentId)
        }
    }

    for _, reQueryId := range reQueryIds {
        res, err := uc.CommentInfo(ctx, reQueryId)
        if err != nil {
            result[reQueryId] = &biz.CommentCount{}
        } else {
            result[reQueryId] = res
        }
    }
    return result, nil
}
