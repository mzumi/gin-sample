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
	connect := os.Getenv("DB_CONNECT")
	db, err = gorm.Open("postgres", connect)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&Article{})
}
