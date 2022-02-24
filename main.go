package main

import (
	"example/sujith/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var student = beans.Student{
// FirstName: "sujith",
// 	LastName:  "vedunuri",
// 	DOB:       "10062001",
// 	Gender:    "male",
// 	City:      "hyderabad",
// 	State:     "Telangana",
// }

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/FirstProject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// result := db.Create(&student)
	// db.Migrator().CreateTable(&student)
	// db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/FirstProject?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established", db)
	}

	// for _, user := range users {
	// 	db.Create(&user)
	// }
	// fmt.Printf("hello")
	r := gin.Default()
	r.GET("/", controllers.GetStudents)
	r.POST("/register", controllers.AddNewStudent)
	r.GET("/:id", controllers.GetStudentById)
	r.Run()
}
