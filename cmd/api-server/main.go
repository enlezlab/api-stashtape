package main

import (
	"fmt"
	"stashtape/handlers"
	"stashtape/types"
)

func main() {
	getCollectionById := handler.GetCollectionItemById
	getCollectionById("123")

	getCollection := handler.GetCollection
	list := getCollection()
	fmt.Println(list)

	addItem := handler.AddItem

	data := types.CollectionItem{
		CollectionId: "ST006",
		Timestamp:    "0000000000",
	}

	addItem(data)

	fmt.Println("run api server")
}
