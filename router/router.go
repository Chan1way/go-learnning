package router

import (
	"go-learnning/handler"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

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
