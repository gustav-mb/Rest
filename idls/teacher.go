package idls

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Teacher struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

var Teachers = []*Teacher{
	{Id: 1, Name: "Rasmus", Rating: 8.5},
	{Id: 2, Name: "Thore", Rating: 9},
	{Id: 3, Name: "Alexandro", Rating: 7.3},
	{Id: 4, Name: "Claus", Rating: 8.6},
	{Id: 5, Name: "Henriette", Rating: 8.2},
}

func GetTeachers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Students)
}

func GetTeacherByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, teacher := findTeacher(id)
	if teacher == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "teacher with given id not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, teacher)
}

func GetTeacherByName(c *gin.Context) {
	name := c.Param("name")

	for _, teacher := range Teachers {
		if teacher.Name == name {
			c.IndentedJSON(http.StatusOK, teacher)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student name not found."})
}

func CreateTeacher(c *gin.Context) {
	var teacher *Teacher

	err := c.BindJSON(teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Teachers = append(Teachers, teacher)
	c.IndentedJSON(http.StatusCreated, teacher)
}

func UpdateTeacher(c *gin.Context) {
	var teacher *Teacher

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, t := findTeacher(id)
	if t == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found."})
		return
	}

	err = c.BindJSON(teacher)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.Id = teacher.Id
	t.Name = teacher.Name
	t.Rating = teacher.Rating

	c.IndentedJSON(http.StatusOK, t)
}

func DeleteTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index, teacher := findTeacher(id)
	if teacher == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "teacher not found"})
		return
	}

	Teachers = removeTeacher(Teachers, index)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "student deleted."})
}

func findTeacher(id int) (int, *Teacher) {
	for i, teacher := range Teachers {
		if teacher.Id == id {
			return i, teacher
		}
	}

	return 0, nil
}

func removeTeacher(slice []*Teacher, s int) []*Teacher {
	return append(slice[:s], slice[s+1:]...)
}
