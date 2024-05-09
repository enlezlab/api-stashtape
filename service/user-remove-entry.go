package service

import (
	"encoding/json"
	"fmt"
	"stashtape/data"
)

func UserRemoveEntry(userId string, uid string) []byte {

	removeEntry := data.RemoveEntry(userId, uid)

	res, err := json.Marshal(removeEntry)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
