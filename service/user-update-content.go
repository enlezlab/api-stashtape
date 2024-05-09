package service

import (
	"stashtape/data"
	"stashtape/model"
)

func UserUpdateContent(newEntry model.User) string {

	entry := data.UpdateEntry("USER_CONTENT", newEntry)

	return entry
}
