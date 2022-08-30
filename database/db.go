package database

import (
	"canvas/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "go_canvas"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return db
}

func Migrate() {
	db := InitDb()
	db.AutoMigrate(&models.Canvas{}, &models.Shape{})
}
