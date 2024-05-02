package handler

import (
	"encoding/json"
	"fmt"
	"stashtape/store"
)

func GetItem(tableName string, valCond string) []byte {

	item := store.GetItem(tableName, valCond)

	res, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
