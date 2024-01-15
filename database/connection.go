package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Init() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open("postgres", fmt.Sprintf("host='%s' port=5432 user=%s dbname='%s' password='%s' sslmode=disable", dbHost, dbUser, dbName, dbPass))

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
