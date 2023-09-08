package service

import (
	"context"

	"github.com/abrarnaim015/belajar-golang-rastful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, req web.CategoryCreateReq) web.CategoryRes
	Update(ctx context.Context, req web.CategoryUpdateReq) web.CategoryRes
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryRes
	FindAll(ctx context.Context) []web.CategoryRes
}

