package router

import (
	"go-learnning/handler"
	"go-learnning/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())
	students := r.Group("/students")
	{
		students.GET("", handler.GetStudents)
		students.GET("/:id", handler.GetStudent)
		students.POST("", handler.CreateStudent)
		students.PUT("/:id", handler.UpdateStudent)
		students.DELETE("/:id", handler.DeleteStudent)
	}

	return r
}
