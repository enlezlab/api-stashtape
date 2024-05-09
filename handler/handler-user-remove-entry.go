package handler

import (
	"net/http"

	"stashtape/service"

	"github.com/go-chi/chi"
)

func HandlerUserRemoveEntry(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "userId")
	entriesId := chi.URLParam(r, "entriesId")

	res := service.UserRemoveEntry(userId, entriesId)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
