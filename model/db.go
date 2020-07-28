package model

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	connect := []string{}

	host := os.Getenv("POSTGRES_HOST")
	if host != "" {
		connect = append(connect, fmt.Sprintf("host=%s", host))
	}

	port := os.Getenv("POSTGRES_PORT")
	if port != "" {
		connect = append(connect, fmt.Sprintf("port=%s", port))
	}

	user := os.Getenv("POSTGRES_USER")
	if user != "" {
		connect = append(connect, fmt.Sprintf("user=%s", user))
	}

	dbName := os.Getenv("POSTGRES_DB_NAME")
	if dbName != "" {
		connect = append(connect, fmt.Sprintf("dbname=%s", dbName))
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password != "" {
		connect = append(connect, fmt.Sprintf("password=%s", password))
	}

	sslmode := os.Getenv("POSTGRES_SSLMODE")
	if sslmode != "" {
		connect = append(connect, fmt.Sprintf("sslmode=%s", sslmode))
	}

	fmt.Println(strings.Join(connect, " "))

	db, err = gorm.Open("postgres", strings.Join(connect, " "))
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&Article{})
}
