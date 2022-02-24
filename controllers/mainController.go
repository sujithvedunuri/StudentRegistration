package controllers

import (
	"example/sujith/beans"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var students []beans.Student = []beans.Student{

	{
		ID:        1,
		FirstName: "sujith",
		LastName:  "vedunuri",
		DOB:       "10062001",
		Gender:    "male",
		City:      "hyderabad",
		State:     "Telangana",
	},
	{
		ID:        2,
		FirstName: "Ricky", LastName: "temp", DOB: "10062001",
		Gender: "male",
		City:   "hyderabad",
		State:  "Telangana"},
	{
		ID:        3,
		FirstName: "Adam", LastName: "temp", DOB: "10062001",
		Gender: "male",
		City:   "hyderabad",
		State:  "Telangana"},
	{
		ID:        4,
		FirstName: "Justin", LastName: "temp", DOB: "10062001",
		Gender: "male",
		City:   "hyderabad",
		State:  "Telangana"},
}

func GetStudents(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK, students)

}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	for _, i := range students {
		if strconv.Itoa(i.ID) == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "student not found",
	})
}
func AddNewStudent(c *gin.Context) {

	var newStudent beans.Student

	if err := c.BindJSON(&newStudent); err != nil {
		log.Fatal("new student added succesfully")
		return

	}
	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)

}
