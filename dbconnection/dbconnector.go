package dbconnection

import (
	"example/sujith/beans"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Dbconnector() {

	dsn := "root:password@tcp(127.0.0.1:3306)/FirstProject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(beans.Student{})
	Db = db
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established", Db)
	}

}
