package main

import (
	"fmt"
	"net/http"
	"stashtape/handler"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.HandleFunc("/user", handler.HandlerNewUser)
	router.HandleFunc("/user/{userId}", handler.HandlerUserContent)
	router.HandleFunc("/content/{userId}", handler.HandlerUserUpdateContent)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}
