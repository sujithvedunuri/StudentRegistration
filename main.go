package main

import (
	"example/sujith/beans"
	"example/sujith/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
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
var err error

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/FirstProject?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db.AutoMigrate(beans.Student{})
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established", Db)
	}

	r := gin.Default()
	r.GET("/", controllers.GetStudents)
	r.POST("/register", func(c *gin.Context) {

		// fmt.Printf("check before")
		var newStudent beans.Student

		if err := c.ShouldBindJSON(&newStudent); err == nil {
			fmt.Printf("obj", newStudent)
		} else {
			fmt.Printf("error here", err)
		}
		fmt.Print(newStudent, "sujith")
		Db.Create(&newStudent)
		controllers.Students = append(controllers.Students, newStudent)
		c.IndentedJSON(http.StatusCreated, newStudent)
	})

	r.GET("/:id", controllers.GetStudentById)
	r.Run()
}
