package main

import (
	"fmt"
	"log"
	"net/http"
	"practice/book-management-system/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server on port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
