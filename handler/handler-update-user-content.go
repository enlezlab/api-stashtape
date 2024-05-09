package handler

import (
	"net/http"
	"stashtape/model"
	"stashtape/service"

	"github.com/go-chi/chi"
)

func HandlerUserUpdateContent(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	entriesMap := []model.Entries{}

	item := model.Entries{
		UID:         "789",
		Title:       "title for 789",
		Description: "desc for 789",
	}

	entriesMap = append(entriesMap, item)

	entries := model.User{
		USER_ID: userId,
		ENTRIES: entriesMap,
	}

	res := service.UserUpdateContent(entries)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}
