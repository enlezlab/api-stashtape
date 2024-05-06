package handler

import (
	"encoding/json"
	"fmt"
	model "stashtape/models"
)

// GetUserContent:
// Get content entries from USER_CONTENT table by USER_ID

func UserGetContent(userId string) []byte {

	getEntry := model.GetEntry("USER_CONTENT", userId)

	res, err := json.Marshal(getEntry)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
