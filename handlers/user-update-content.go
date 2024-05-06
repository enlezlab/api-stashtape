package handler

import (
	"fmt"
	model "stashtape/models"
	"stashtape/types"
)

func UserUpdateContent(newEntry types.User) {

	entry := model.UpdateEntry("USER_CONTENT", newEntry)
	fmt.Println(entry)

}
