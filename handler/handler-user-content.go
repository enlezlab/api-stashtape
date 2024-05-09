package handler

import (
	"net/http"
	"stashtape/service"

	"github.com/go-chi/chi"
)

func HandlerUserContent(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userGetContent := service.UserGetContent
	list := userGetContent(userId)
	w.Header().Set("Content-Type", "application/json")
	w.Write(list)
}
