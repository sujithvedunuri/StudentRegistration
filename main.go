package main

import (
	"example/sujith/beans"
	"example/sujith/controllers"
	"example/sujith/dbconnection"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var stud = beans.Student{
	FirstName: "sujith",
	LastName:  "vedunuri",
	DOB:       "10062001",
	Gender:    "male",
	City:      "hyderabad",
	State:     "Telangana",
	// CreatedAt:  time.Now(),
	// ApprovedAt: time.Now(),
	// RejectedAt: time.Now(),
}
var Db *gorm.DB

func main() {

	r := gin.Default()
	dbconnection.Dbconnector()
	r.GET("/", controllers.GetStudents)
	r.POST("/register", controllers.RegisterStudents)
	r.GET("/:id", controllers.GetStudentById)
	r.Run()
}
