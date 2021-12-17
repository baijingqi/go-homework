package data

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "user/internal/biz"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, g *biz.User) error {
	return nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, g *biz.User) error {
	return nil
}
