package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 记录请求耗时
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // 执行后续handler

		duration := time.Since(start)
		fmt.Printf("[%s] %s %s 耗时: %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			duration,
		)
	}
}

// 捕获 panic，防止服务器崩溃
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v", err)
				c.JSON(500, gin.H{"error": "服务器内部错误"})
				c.Abort()
			}
		}()
		c.Next()
	}
}
