package handler

import (
	"net/http"
	"stashtape/model"
	"stashtape/service"
)

func HandlerNewUser(w http.ResponseWriter, r *http.Request) {

	entries := model.Entries{
		UID:         "123123",
		Title:       "title new",
		Description: "desc new",
	}

	res := service.NewUser(entries)

	// res := "{\"test\": \"testval\"}"

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}
