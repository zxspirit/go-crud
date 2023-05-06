package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const DbUsername = "root"
const DbPassword = "Zhengxu1./"
const DbName = "go"
const DbHost = "localhost"
const DbPort = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDb()

	return Db
}

func connectDb() *gorm.DB {
	var err error
	dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?" + "parseTime=true&loc=Local"
	log.Println("dsn : ", dsn)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
		return nil
	}
	return open

}
