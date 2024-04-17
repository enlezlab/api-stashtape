package main

import (
	"fmt"
	"stashtape/handlers"
	"stashtape/types"
)

func main() {

	fmt.Println("===get all collection from table example===")
	getCollection := handler.GetCollection
	list := getCollection("collection")
	fmt.Println(list)

	fmt.Println("===end of get all collection from table example===")

	fmt.Println("")

	fmt.Println("===get item from table example===")

	getItem := handler.GetItem
	item := getItem("collection", "ST012")
	fmt.Println(item)

	fmt.Println("===end of get item from table example===")

	fmt.Println("")

	fmt.Println("===add Item to table example===")
	addItem := handler.AddItem

	data := types.CollectionItem{
		CollectionId: "ST012",
		Timestamp:    "7777777777",
	}

	res := addItem("collection", data)
	fmt.Println(res)
	fmt.Println("===end of add item to table example===")

}
