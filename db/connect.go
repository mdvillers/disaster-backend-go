package db

import (
	"fmt"

	"github.com/mdvillers/disaster-backend-go/config"
	"github.com/mdvillers/disaster-backend-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB instance
var (
	DB *gorm.DB
)

//Connect to database
func Connect() {
	dsn := config.DBURI
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	db.AutoMigrate(&models.Admin{})
	DB = db
}
