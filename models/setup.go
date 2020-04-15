package models

import (
	"github.com/jinzhu/gorm"

	// import mysql connector for golang
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/torkashvand/goshortener/config"
	"github.com/torkashvand/goshortener/log"

	// import sqlite connector for golang
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//InitDB setup connection and create inital schema for our models
func InitDB() *gorm.DB {
	config := config.Config()
	connectionString := config.GetString("DB_CONNECTION_STRING")
	dbDriver := config.GetString("DB_DRIVER")

	db, err := gorm.Open(dbDriver, connectionString)

	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&Link{})

	return db
}
