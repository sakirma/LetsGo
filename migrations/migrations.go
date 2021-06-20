package main

import (
	"LetsGooooo/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/LetsGo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Failed gorm.Open: %v\n", err)
	}

	DB = db
	creatTables(models.User{})
}

func creatTables(dst ...interface{}) {
	err := DB.AutoMigrate(dst)

	if err != nil {
		log.Panicf("Failed to create Table", err)
	}
}
