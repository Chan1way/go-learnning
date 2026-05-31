package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func startGinServer() {
	initDB()
	r := gin.Default()

	// 路由1：返回文字
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	// // 路由2：返回JSON
	// r.GET("/student", func(c *gin.Context) {
	// 	s := Student{Name: "小明", Score: 95}
	// 	c.JSON(http.StatusOK, s)
	// })

	setupDBCRUD(r)
	r.Run(":8080")
}
