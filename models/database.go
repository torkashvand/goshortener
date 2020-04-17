package models

import (
	"os"

	"github.com/jinzhu/gorm"

	// import mysql connector for golang
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/torkashvand/goshortener/config"
	"github.com/torkashvand/goshortener/log"

	// import sqlite connector for golang
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


// DatabaseProvide repersent abestraction around database
type DatabaseProvide interface {
	Open()
	Close()
	AutoMigrate()
	SetDB(db *gorm.DB)
	GetDB() (db *gorm.DB)
}

// Database base struct
type Database struct {
	db *gorm.DB
}

// GetDB is the getter
func (database *Database) GetDB() *gorm.DB {
	return database.db
}

// SetDB is the setter
func (database *Database) SetDB(db *gorm.DB) {
	database.db = db
}

// AutoMigrate is responsible for creating database from models
func (database *Database) AutoMigrate() {
	database.db.AutoMigrate(&Link{})
}

// Close the connection to databse
func (database *Database) Close() {
	database.db.Close()
}

// MysqlDB reperesent concrete struct for Mysql database
type MysqlDB struct {
	Database
}

// Open opens a connection to database
func (mysql *MysqlDB) Open() {
	config := config.Config()
	connectionString := config.GetString("DB_CONNECTION_STRING")
	dbDriver := config.GetString("DB_DRIVER")

	db, err := gorm.Open(dbDriver, connectionString)

	if err != nil {
		log.Panic(err)
	}

	mysql.SetDB(db)
}

// SQLite reperesent concrete struct for SQLite database
type SQLite struct {
	Database
}

// Open opens a connection to database
func (sqlite *SQLite) Open() {
	config := config.Config()
	connectionString := config.GetString("DB_CONNECTION_STRING")
	dbDriver := config.GetString("DB_DRIVER")

	db, err := gorm.Open(dbDriver, connectionString)

	if err != nil {
		log.Panic(err)
	}

	sqlite.SetDB(db)
}

// Close the connection to databse
func (sqlite *SQLite) Close() {
	sqlite.db.Close()

	os.Remove(config.Config().GetString("DB_CONNECTION_STRING"))
}
