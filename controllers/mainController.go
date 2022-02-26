package controllers

import (
	"example/sujith/beans"
	"example/sujith/dbconnection"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

var Students []beans.Student = []beans.Student{

	// {
	// 	ID:        1,
	// 	FirstName: "sujith",
	// 	LastName:  "vedunuri",
	// 	DOB:       "10062001",
	// 	Gender:    "male",
	// 	City:      "hyderabad",
	// 	State:     "Telangana",
	// },
	// {
	// 	ID:        2,
	// 	FirstName: "Ricky", LastName: "temp", DOB: "10062001",
	// 	Gender: "male",
	// 	City:   "hyderabad",
	// 	State:  "Telangana"},
	// {
	// 	ID:        3,
	// 	FirstName: "Adam", LastName: "temp", DOB: "10062001",
	// 	Gender: "male",
	// 	City:   "hyderabad",
	// 	State:  "Telangana"},
	// {
	// 	ID:        4,
	// 	FirstName: "Justin", LastName: "temp", DOB: "10062001",
	// 	Gender: "male",
	// 	City:   "hyderabad",
	// 	State:  "Telangana"},
}

func RegisterStudents(c *gin.Context) {

	// fmt.Printf("check before")
	var newStudent beans.Student

	if err := c.ShouldBindJSON(&newStudent); err == nil {
		fmt.Printf("obj", newStudent)
	} else {
		fmt.Printf("error here", err)
	}
	fmt.Print(newStudent, "sujith")
	dbconnection.Db.Create(&newStudent)
	Students = append(Students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func GetStudents(c *gin.Context) {
	// var studes []beans.Student
	var studata []beans.Student
	dbconnection.Db.Find(&studata)
	fmt.Println(studata)

	// fmt.Println("printfing body", c.Request.Body)
	// fmt.Println("dujith", c.Request.Response)

	// if err := c.ShouldBindJSON(&studes); err == nil {
	// 	fmt.Println("list", studes)
	// } else {
	// 	fmt.Println("error", err)
	// }

	// Students tudes
	c.IndentedJSON(
		http.StatusOK, studata)

}

func GetStudentById(c *gin.Context) {
	var id = c.Param("id")
	var query []beans.Student
	dbconnection.Db.Where("id", id).Find(&query)
	fmt.Println("the searched query is: ", query)

	c.IndentedJSON(
		http.StatusOK, query)
}
