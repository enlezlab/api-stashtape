package handlerfunc

import (
	"net/http"
	handler "stashtape/handlers"

	"github.com/go-chi/chi"
)

func HandlerUserContent(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userGetContent := handler.UserGetContent
	list := userGetContent(userId)
	w.Header().Set("Content-Type", "application/json")
	w.Write(list)
}
