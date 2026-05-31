package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupDBCRUD(r *gin.Engine) {
	// 查：获取所有学生
	r.GET("/students", func(c *gin.Context) {
		var students []Student
		DB.Find(&students)
		c.JSON(http.StatusOK, students)
	})

	// 查：获取单个学生
	r.GET("/students/:id", func(c *gin.Context) {
		var student Student
		result := DB.First(&student, c.Param("id"))
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
			return
		}
		c.JSON(http.StatusOK, student)
	})

	// 增：新增学生
	r.POST("/students", func(c *gin.Context) {
		var student Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}
		result := DB.Create(&student)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "新增成功", "student": student})
	})

	// 改： 修改学生
	r.PUT("/students/:id", func(c *gin.Context) {
		var student Student
		result := DB.First(&student, c.Param("id"))
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
			return
		}
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}
		DB.Save(&student)
		c.JSON(http.StatusOK, gin.H{"message": "修改成功", "student": student})
	})

	// 删：删除学生
	r.DELETE("/students/:id", func(c *gin.Context) {
		result := DB.Delete(&Student{}, c.Param("id"))
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
	})
}

//
