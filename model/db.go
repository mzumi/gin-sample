package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	connect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE"))
	db, err = gorm.Open("postgres", connect)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&Article{})
}
