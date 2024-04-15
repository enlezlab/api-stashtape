package main

import (
	"fmt"
	"stashtape/handlers"
	// "stashtape/types"
)

func main() {

	getCollection := handler.GetCollection
	list := getCollection()
	fmt.Println(list)

	getCollectionItem := handler.GetCollectionItem
	getCollectionItem()

	// addCollectionItem := handler.AddCollectionItem

	// data := types.CollectionItem{
	// 	CollectionId: "ST007",
	// 	Timestamp:    "0000000000",
	// }

	// addCollectionItem(data)

	fmt.Println("run api server")
}
