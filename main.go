package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/abrarnaim015/belajar-golang-rastful-api/app"
	"github.com/abrarnaim015/belajar-golang-rastful-api/controller"
	"github.com/abrarnaim015/belajar-golang-rastful-api/helper"
	"github.com/abrarnaim015/belajar-golang-rastful-api/middleware"
	"github.com/abrarnaim015/belajar-golang-rastful-api/repository"
	"github.com/abrarnaim015/belajar-golang-rastful-api/service"
	"github.com/go-playground/validator"
)

func main()  {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server {
		Addr: ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("Server Has Running At http://localhost" + server.Addr + " 🚀")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}