package controller

import (
	"net/http"
	"strconv"

	"github.com/abrarnaim015/belajar-golang-rastful-api/helper"
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/web"
	"github.com/abrarnaim015/belajar-golang-rastful-api/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	categoryCreateReq := web.CategoryCreateReq {}

	helper.ReadFromReqBody(r, &categoryCreateReq)

	categoryRes := controller.CategoryService.Create(r.Context(), categoryCreateReq)
	w.WriteHeader(http.StatusCreated)
	
	webRes := web.WebRes {
		Code: http.StatusCreated,
		Status: "Succsess Create New Category",
		Data: categoryRes,
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	categoryUpdateReq := web.CategoryUpdateReq {}

	helper.ReadFromReqBody(r, &categoryUpdateReq)

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateReq.Id = id

	categoryRes := controller.CategoryService.Update(r.Context(), categoryUpdateReq)
	webRes := web.WebRes {
		Code: http.StatusOK,
		Status: "Succsess Update Category",
		Data: categoryRes,
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)

	webRes := web.WebRes {
		Code: http.StatusOK,
		Status: "Succsess Delete Category Id:" + categoryId,
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryRes := controller.CategoryService.FindById(r.Context(),id)

	webRes := web.WebRes {
		Code: http.StatusOK,
		Status: "Succsess Find Category Id:" + categoryId,
		Data: categoryRes,
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	categoryReses := controller.CategoryService.FindAll(r.Context())

	webRes := web.WebRes {
		Code: http.StatusOK,
		Status: "Succsess Get All Category",
		Data: categoryReses,
	}

	helper.WriteToResBody(w, webRes)
}