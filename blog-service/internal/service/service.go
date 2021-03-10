package service

import (
	"context"

	"github.com/yann0917/go-tour-book/blog-service/global"
	"github.com/yann0917/go-tour-book/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}
