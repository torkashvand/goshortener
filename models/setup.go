package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// import mysql connector for golang
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/torkashvand/goshortener/config"

	// import sqlite connector for golang
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//SetupModels setup connection and create inital schema for our models
func SetupModels() *gorm.DB {
	config := config.Config()
	connectionString := config.GetString("DB_CONNECTION_STRING")
	fmt.Println("DB_CONNECTION_STRING: " + connectionString)
	dbDriver := config.GetString("DB_DRIVER")
	
	db, err := gorm.Open(dbDriver, connectionString)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Link{})

	return db
}
