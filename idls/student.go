package idls

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Enrollment     string  `json:"enrollment"`
	CourseWorkLoad float64 `json:"courseworkload"`
}

var Students = []*Student{
	{Id: 1, Name: "John", Enrollment: "Active", CourseWorkLoad: 30},
	{Id: 2, Name: "Alice", Enrollment: "Dropout", CourseWorkLoad: 0},
	{Id: 3, Name: "Bob", Enrollment: "Graduated", CourseWorkLoad: 0},
	{Id: 4, Name: "Peter", Enrollment: "Active", CourseWorkLoad: 25},
	{Id: 5, Name: "Jane", Enrollment: "New", CourseWorkLoad: 0},
}

func GetStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Students)
}

func GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, student := findStudent(id)
	if student == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student with given id not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func GetStudentByName(c *gin.Context) {
	name := c.Param("name")

	for _, student := range Students {
		if student.Name == name {
			c.IndentedJSON(http.StatusOK, student)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student name not found."})
}

func CreateStudent(c *gin.Context) {
	var student Student

	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Students = append(Students, &student)
	c.IndentedJSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	var student *Student

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, s := findStudent(id)
	if s == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found."})
		return
	}

	err = c.BindJSON(student)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.Id = student.Id
	s.Name = student.Name
	s.Enrollment = student.Enrollment
	s.CourseWorkLoad = student.CourseWorkLoad

	c.IndentedJSON(http.StatusOK, s)
}

func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index, student := findStudent(id)
	if student == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
		return
	}

	Students = removeStudent(Students, index)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "student deleted."})
}

func findStudent(id int) (int, *Student) {
	for i, student := range Students {
		if student.Id == id {
			return i, student
		}
	}

	return 0, nil
}

func removeStudent(slice []*Student, s int) []*Student {
	return append(slice[:s], slice[s+1:]...)
}
