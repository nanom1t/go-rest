package models

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

// Initialize database
func Initialize() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)

	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	database = db

	return database, nil
}

// Get database
func GetDatabase() (*gorm.DB) {
	return database
}