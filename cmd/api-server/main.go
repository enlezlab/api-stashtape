package main

import (
	"fmt"
	"net/http"
	handlerfunc "stashtape/handler-func"

	"github.com/go-chi/chi"
	// "stashtape/types"
	// "encoding/json"
)

func main() {
	router := chi.NewRouter()
	router.HandleFunc("/{userId}", handlerfunc.HandlerUserContent)
	router.HandleFunc("/content/{userId}", handlerfunc.HandlerUserUpdateContent)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)

	// fmt.Println("===delete item===")

	// deleteItem := handler.DeleteItem
	// deleteItem("collection", "ST002")

	// fmt.Println("===end of delete item===")

	// fmt.Println("===get all collection from table example===")
	// getCollection := handler.GetCollection
	// list := getCollection("collection")
	// fmt.Println(list)

	// fmt.Println("===end of get all collection from table example===")

	// fmt.Println("")

	// fmt.Println("===get item from table example===")

	// getItem := handler.GetItem
	// item := getItem("collection", "ST011")
	// fmt.Println(item)

	// fmt.Println("===end of get item from table example===")

	// fmt.Println("===get entry from table example===")

	// getEntry := model.GetEntry
	// item := getEntry("USER", "user_id_here")
	// fmt.Println(item)

	// fmt.Println("===end of get entry from table example===")

	// fmt.Println("")

	// fmt.Println("===add Item to table example===")

	// entry := []types.Entries{}

	// item := types.Entries{
	// 	UID:         "new entry_hash_here",
	// 	Title:       "new Entry title here",
	// 	Description: "new Entry description here",
	// }

	// entry = append(entry, item)

	// userId := util.UserIdGen()

	// data := types.User{
	// 	USER_ID: userId,
	// 	ENTRIES: entry,
	// }

	// newEntry := model.NewEntry
	// res := newEntry("USER_CONTENT", data)
	// fmt.Println(res)
	// fmt.Println("===end of add item to table example===")

	// fmt.Println("===update user entry to table example===")
	// updateEntry := model.UpdateEntry

	// var colList []types.Entries

	// item := types.Entries{
	// 	UID:         "uid here",
	// 	Title:       "entry title here",
	// 	Description: "entry description here",
	// }

	// colList = append(colList, item)

	// data := types.User{
	// 	USER_ID: "f7652dad-8b6c-47d4-b185-c6db76782812",
	// 	ENTRIES: colList,
	// }

	// res := updateEntry("USER_CONTENT", data)
	// fmt.Println(res)
	// fmt.Println("===end of update user entry to table example===")

	// fmt.Println("===update Item to table example===")
	// updateItem := handler.UpdateItem

	// var colList []types.CollectionItemList

	// item := types.CollectionItemList{
	// 	Title:       "3 testing adding new item to exsiting user",
	// 	Description: "3 desc for testing add new item to existing user",
	// }

	// colList = append(colList, item)

	// data := types.CollectionItem{
	// 	CollectionId: "ST015",
	// 	Timestamp:    "6666666666",
	// 	List:         colList,
	// }

	// res := updateItem("collection", data)
	// fmt.Println(res)
	// fmt.Println("===end of update item to table example===")

}
