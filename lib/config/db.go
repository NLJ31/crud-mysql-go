package config

import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-gorm?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connection successful")
	db = d
}

func Database() *gorm.DB {
	return db
}

