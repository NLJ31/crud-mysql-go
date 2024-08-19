package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NLJ31/crud-mysql-go/lib/config"
	"github.com/NLJ31/crud-mysql-go/lib/models"
	"github.com/NLJ31/crud-mysql-go/lib/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB


func connectDB() {
	config.Connect();
	db = config.Database();
}

func CreateHero(res http.ResponseWriter, req *http.Request) {
	var hero models.Hero;

	utils.ReqBodyParse(req, &hero);
	
	connectDB();
	
	err := hero.CreateUser(db);

	if err != nil {
		utils.Failure(res, 404, &utils.FailureBody{
			Status: 404, 
			ErrorMessage: "Database Error", 
			Message: "Error WHen Creating",
		});
		return
	}

	utils.Success(res, &utils.SuccessBody{
		Version: "v1", 
		Status: 200, 
		Message: "Success", 
		Data: hero,
	});
}

func GetAllHero(res http.ResponseWriter, req *http.Request) {
	var hero []models.Hero;
	// var model = &models.Hero{};

	connectDB();
	
	err := models.GetAllHeros(db, &hero);

	if err != nil {
		print(err);
		utils.Failure(res, 404, &utils.FailureBody{
			Status: 404, 
			ErrorMessage: "Database Error", 
			Message: "Error When Fetching Data",
		});
		return
	}

	utils.Success(res, &utils.SuccessBody{
		Version: "v1", 
		Status: 200, 
		Message: "Success", 
		Data: hero,
	});
}

func DeleteHero(res http.ResponseWriter, req *http.Request) {
	
	var hero models.Hero;

	utils.ParamsParse(req, &hero)

	fmt.Println(hero.ID);
	
	connectDB();

	err := hero.DeleteHero(db);

	if err != nil {
		print(err);
		utils.Failure(res, 404, &utils.FailureBody{
			Status: 404, 
			ErrorMessage: "Database Error", 
			Message: "Error When Deleting Data",
		});
		return
	}

	utils.Success(res, &utils.SuccessBody{
		Version: "v1", 
		Status: 200, 
		Message: "Success", 
		Data: hero,
	});

}


func HeroDetail(res http.ResponseWriter, req *http.Request) {
	
	var hero models.Hero;
	
	connectDB();
    params := mux.Vars(req);

	id, errStr := strconv.Atoi(params["id"]);

	if (errStr != nil) {
		return
	}
	

	err := hero.GetDetail(db, id);

	if err != nil {
		print(err);
		utils.Failure(res, 404, &utils.FailureBody{
			Status: 404, 
			ErrorMessage: "Database Error", 
			Message: "Error When Deleting Data",
		});
		return
	}

	utils.Success(res, &utils.SuccessBody{
		Version: "v1", 
		Status: 200, 
		Message: "Success", 
		Data: hero,
	});

}

