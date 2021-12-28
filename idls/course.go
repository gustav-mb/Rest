package idls

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Course struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Teacher *Teacher `json:"teacher"`
	Score   float64  `json:"score"`
}

var Courses = []*Course{
	{Id: 1, Name: "BDSA", Teacher: Teachers[0], Score: 7.0},
	{Id: 2, Name: "ALGO", Teacher: Teachers[1], Score: 9.0},
	{Id: 3, Name: "DMAT", Teacher: Teachers[2], Score: 9.5},
	{Id: 4, Name: "GRPRO", Teacher: Teachers[3], Score: 10.0},
	{Id: 5, Name: "PRKOM", Teacher: Teachers[4], Score: 10.0},
}

func GetCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Courses)
}

func GetCourseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, course := findCourse(id)
	if course == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course id not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, course)
}

func GetCourseByName(c *gin.Context) {
	name := c.Param("name")

	for _, course := range Courses {
		if course.Name == name {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course name not found."})
}

func CreateCourse(c *gin.Context) {
	var course *Course

	err := c.BindJSON(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Courses = append(Courses, course)
	c.IndentedJSON(http.StatusCreated, course)
}

func UpdateCourse(c *gin.Context) {
	var course *Course

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, co := findCourse(id)
	if co == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found."})
		return
	}

	err = c.BindJSON(course)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	co.Id = course.Id
	co.Name = course.Name
	co.Teacher = course.Teacher
	co.Score = course.Score

	c.IndentedJSON(http.StatusOK, co)
}

func DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index, course := findCourse(id)
	if course == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
		return
	}

	Courses = removeCourse(Courses, index)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "course deleted."})
}

func findCourse(id int) (int, *Course) {
	for i, course := range Courses {
		if course.Id == id {
			return i, course
		}
	}

	return 0, nil
}

func removeCourse(slice []*Course, s int) []*Course {
	return append(slice[:s], slice[s+1:]...)
}
