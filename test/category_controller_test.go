package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/abrarnaim015/belajar-golang-rastful-api/app"
	"github.com/abrarnaim015/belajar-golang-rastful-api/controller"
	"github.com/abrarnaim015/belajar-golang-rastful-api/helper"
	"github.com/abrarnaim015/belajar-golang-rastful-api/middleware"
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/domain"
	"github.com/abrarnaim015/belajar-golang-rastful-api/repository"
	"github.com/abrarnaim015/belajar-golang-rastful-api/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_resful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler  {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB)  {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)
	
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"name": "Gadget"}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", reqBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	resBody := map[string]interface{} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Equal(t, http.StatusCreated, int(resBody["code"].(float64)))
	assert.Equal(t, "Succsess Create New Category", resBody["status"])
	assert.Equal(t, "Gadget", resBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	router := setupRouter(db)

	reqBody := strings.NewReader(`{"name": ""}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", reqBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	msgErr := "Key: 'CategoryCreateReq.Name' Error:Field validation for 'Name' failed on the 'required' tag"

	assert.Equal(t, http.StatusBadRequest, int(resBody["code"].(float64)))
	assert.Equal(t, "Bad Request", resBody["status"])
	assert.Equal(t, msgErr, resBody["data"])
}

func TestUpdateCategorySuccess(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	tx, _ := db.Begin()
	category := repository.NewCategoryRepository().Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	reqBody := strings.NewReader(`{"name": "Gadget Update"}`)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), reqBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, "Succsess Update Category", resBody["status"])
	assert.Equal(t, category.Id, int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget Update", resBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	tx, _ := db.Begin()
	category := repository.NewCategoryRepository().Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	id := category.Id + 1

	router := setupRouter(db)

	reqBody := strings.NewReader(`{"name": "Gadget Update"}`)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+ strconv.Itoa(id), reqBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusNotFound, int(resBody["code"].(float64)))
	assert.Equal(t, "Not Found", resBody["status"])
	assert.Equal(t, "Category is not found", resBody["data"])
}

func TestGetCategorySuccess(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	tx, _ := db.Begin()
	category := repository.NewCategoryRepository().Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, "Succsess Find Category Id:"+strconv.Itoa(category.Id), resBody["status"])
	assert.Equal(t, category.Id, int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", resBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	router := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusNotFound, int(resBody["code"].(float64)))
	assert.Equal(t, "Not Found", resBody["status"])
	assert.Equal(t, "Category is not found", resBody["data"])
}

func TestDeleteCategorySuccess(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	tx, _ := db.Begin()
	category := repository.NewCategoryRepository().Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, "Succsess Delete Category Id:"+strconv.Itoa(category.Id), resBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	router := setupRouter(db)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusNotFound, int(resBody["code"].(float64)))
	assert.Equal(t, "Not Found", resBody["status"])
	assert.Equal(t, "Category is not found", resBody["data"])
}

func TestListCategorySuccess(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	tx, _ := db.Begin()
	category := repository.NewCategoryRepository().Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	categories := resBody["data"].([]interface{})
	categoryRes := categories[0].(map[string]interface{})

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, "Succsess Get All Category", resBody["status"])
	assert.Equal(t, category.Id, int(categoryRes["id"].(float64)))
	assert.Equal(t, category.Name, categoryRes["name"])
}

func TestUnauthorized(t *testing.T)  {
	db := setupTestDB()
	defer db.Close()
	truncateCategory(db)

	router := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "SALAH")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	resBody := map[string]interface {} {}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusUnauthorized, int(resBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", resBody["status"])
	assert.Equal(t, nil, resBody["data"])
}