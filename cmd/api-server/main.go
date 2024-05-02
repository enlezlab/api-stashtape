package main

import (
	"fmt"
	"net/http"
	handler "stashtape/handlers"

	"github.com/go-chi/chi"
	// "stashtape/types"
	// "encoding/json"
)

func handlerCollectionItem(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "collectionId")

	getItem := handler.GetItem
	item := getItem("collection", id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(item)
}

func handlerCollection(w http.ResponseWriter, r *http.Request) {
	getCollection := handler.GetCollection
	list := getCollection("collection")
	w.Header().Set("Content-Type", "application/json")
	w.Write(list)
}

func main() {
	router := chi.NewRouter()
	router.HandleFunc("/collection/{collectionId}", handlerCollectionItem)
	router.HandleFunc("/collection", handlerCollection)
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

	// fmt.Println("")

	// fmt.Println("===add Item to table example===")
	// addItem := handler.AddItem

	// data := types.CollectionItem{
	// 	CollectionId: "ST015",
	// 	Timestamp:    "6666666666",
	// }

	// res := addItem("collection", data)
	// fmt.Println(res)
	// fmt.Println("===end of add item to table example===")

}
