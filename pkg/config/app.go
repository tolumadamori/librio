package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Initializes the connection to the database. Gets db credentials from env variable set as user: "mysqluser" and password: "mysqluserpassword"
func ConnectDB() (db *gorm.DB) {
	user := os.Getenv("mysqluser")
	password := os.Getenv("mysqluserpassword")
	mysqlcredentials := string(user) + ":" + string(password) + "@tcp(127.0.0.1:3306)/librio?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlcredentials), &gorm.Config{})
	if err != nil {
		fmt.Printf("There was an error connecting to the DB: %v", err)
	}
	return db
}
