package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"idls"
	"log"
	"net"
	"net/http"
	"strconv"
)

type Server struct {
	ipAddress *net.TCPAddr
	router    *gin.Engine
}

func main() {
	var address = flag.String("address", "localhost", "The address of the server.")
	var port = flag.Int("port", 8080, "The port of the server.")
	flag.Parse()

	server := NewServer(*address, *port)
	server.Init()
}

func (s *Server) Init() {
	// Test
	s.router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	// Students
	s.router.GET("/students", idls.GetStudents)
	s.router.GET("/students/id/:id", idls.GetStudentByID)
	s.router.GET("/students/name/:name", idls.GetStudentByName)
	s.router.POST("/students", idls.CreateStudent)
	s.router.PUT("/students/id/:id", idls.UpdateStudent)
	s.router.DELETE("/students/id/:id", idls.DeleteStudent)

	// Courses
	s.router.GET("/courses", idls.GetCourses)
	s.router.GET("/courses/id/:id", idls.GetCourseByID)
	s.router.GET("/courses/name/:name", idls.GetCourseByName)
	s.router.POST("/courses", idls.CreateCourse)
	s.router.PUT("/courses/id/:id", idls.UpdateCourse)
	s.router.DELETE("/courses/id/:id", idls.DeleteCourse)

	// Teachers
	s.router.GET("/teachers", idls.GetTeachers)
	s.router.GET("/teachers/id/:id", idls.GetTeacherByID)
	s.router.GET("/teachers/name/:name", idls.GetTeacherByName)
	s.router.POST("/teachers", idls.CreateTeacher)
	s.router.PUT("/teachers/id/:id", idls.UpdateTeacher)
	s.router.DELETE("/teachers/id/:id", idls.DeleteTeacher)

	// Start server
	err := s.router.Run(s.ipAddress.String())
	if err != nil {
		log.Fatalf("Could not start server. :: %v", err)
	}
}

func NewServer(address string, port int) *Server {
	ipAddress, err := net.ResolveTCPAddr("tcp", address+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Could not resolve ip address %v:%v :: %v", address, port, err)
	}

	return &Server{
		ipAddress: ipAddress,
		router:    gin.Default(),
	}
}
