package exception

import (
	"net/http"

	"github.com/abrarnaim015/belajar-golang-rastful-api/helper"
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/web"
	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface {})  {
	if notFoundError (w, r, err) {
		return
	}

	if validateError (w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface {})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webRes := web.WebRes {
		Code: http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data: err,
	}

	helper.WriteToResBody(w, webRes)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface {}) bool  {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webRes := web.WebRes {
			Code: http.StatusNotFound,
			Status: "Not Found",
			Data: exception.Error,
		}

		helper.WriteToResBody(w, webRes)
		return true
	} else {
		return false
	}
}

func validateError(w http.ResponseWriter, r *http.Request, err interface {}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webRes := web.WebRes {
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Data: exception.Error(),
		}

		helper.WriteToResBody(w, webRes)
		return true
	} else {
		return false
	}
}