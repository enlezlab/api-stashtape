package handlerfunc

import (
	"net/http"
	handler "stashtape/handlers"
	"stashtape/types"

	"github.com/go-chi/chi"
)

func HandlerUserUpdateContent(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	var entries []types.Entries

	item := types.Entries{
		UID:         "789",
		Title:       "title for 789",
		Description: "desc for 789",
	}

	entries = append(entries, item)

	entry := types.User{
		USER_ID: userId,
		ENTRIES: entries,
	}

	userUpdateContent := handler.UserUpdateContent
	userUpdateContent(entry)
	// list := userUpdateContent(userId, entry)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(list)
}
