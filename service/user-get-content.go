package service

import (
	"encoding/json"
	"fmt"
	"stashtape/data"
)

// GetUserContent:
// Get content entries from USER_CONTENT table by USER_ID

func UserGetContent(userId string) []byte {

	getEntry := data.GetEntry("USER_CONTENT", userId)

	res, err := json.Marshal(getEntry)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
