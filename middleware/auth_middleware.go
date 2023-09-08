package middleware

import (
	"net/http"

	"github.com/abrarnaim015/belajar-golang-rastful-api/helper"
	"github.com/abrarnaim015/belajar-golang-rastful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware  {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webRes := web.WebRes {
			Code: http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResBody(w, webRes)
	}

}