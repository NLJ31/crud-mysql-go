package main

import (
	"log"
	"net/http"

	"github.com/NLJ31/crud-mysql-go/lib/routes"
	"github.com/gorilla/mux"
)

// var db *gorm.DB


func main() {

	r := mux.NewRouter()
	routes.HeroRoutes(r)
	http.Handle("/", r)
	log.Println("Listening in post 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}