package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
)

type Praise struct {
    Hello string
}

type PraiseRepo interface {
    AddPraise(context.Context, *Praise) error
    CancelPraise(context.Context, *Praise) error
    IsPraise(context.Context, *Praise) error
}

type PraiseUsecase struct {
    repo PraiseRepo
    log  *log.Helper
}

func NewPraiseUsecase(repo PraiseRepo, logger log.Logger) *PraiseUsecase {
    return &PraiseUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PraiseUsecase) Create(ctx context.Context, g *Praise) error {
    return uc.repo.AddPraise(ctx, g)
}

func (uc *PraiseUsecase) Update(ctx context.Context, g *Praise) error {
    return uc.repo.CancelPraise(ctx, g)
}
