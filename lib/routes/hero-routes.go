package routes

import (
	"github.com/NLJ31/crud-mysql-go/lib/controllers"
	"github.com/gorilla/mux"
)

var HeroRoutes = func(router *mux.Router) {
	router.HandleFunc("/Hero/create", controllers.CreateHero).Methods("POST");
	router.HandleFunc("/Hero/get-all", controllers.GetAllHero).Methods("GET");
	router.HandleFunc("/Hero/{id}/detail", controllers.HeroDetail).Methods("GET");
	router.HandleFunc("/Hero/{id}/delete", controllers.DeleteHero).Methods("DELETE");

}