package service

import (
	"context"
	"database/sql"

	"github.com/abrarnaim015/belajar-golang-rastful-api/exception"
	"github.com/abrarnaim015/belajar-golang-rastful-api/helper"
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/domain"
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/web"
	"github.com/abrarnaim015/belajar-golang-rastful-api/repository"
	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(categoryrepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryrepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, req web.CategoryCreateReq) web.CategoryRes  {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category {
		Name: req.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryRes(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, req web.CategoryUpdateReq) web.CategoryRes  {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = req.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryRes(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int)  {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryRes  {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryRes(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryRes  {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResult(categories)
}