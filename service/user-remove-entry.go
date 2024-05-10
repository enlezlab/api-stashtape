package service

import (
	"encoding/json"
	"fmt"
	"stashtape/data"
	"stashtape/model"
)

func UserRemoveEntry(userId string, uid string) []byte {

	entries := data.GetEntry("USER_CONTENT", userId)
	originEntries := entries[0].ENTRIES

	newEntries := []model.Entries{}

	for _, item := range originEntries {
		if item.UID != uid {
			newEntries = append(newEntries, item)
		}
	}

	fmt.Println(newEntries)

	result := data.NewEntry("USER_CONTENT", model.User{
		USER_ID: userId,
		ENTRIES: newEntries,
	})

	res, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
