package db

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/sqlserver"
)

var DB *gorm.DB
var DSN = "sqlserver://sa:Eco123456(192.168.27.12:1433)/trazabilidad?charset=utf8mb4&parseTime=True&loc=Local"

func DBConnection(){
	var error error
	DB, error = gorm.Open(sqlserver.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database is connected")
	}
}