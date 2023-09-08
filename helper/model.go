package helper

import (
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/domain"
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/web"
)

func ToCategoryRes(category domain.Category) web.CategoryRes  {
	return web.CategoryRes{
		Id: category.Id,
		Name: category.Name,
	}
}

func ToCategoryResult(categories []domain.Category) []web.CategoryRes  {
	categoryResult := []web.CategoryRes {}

	for _, category := range categories {
		categoryResult = append(categoryResult, ToCategoryRes(category))
	}

	return categoryResult
}