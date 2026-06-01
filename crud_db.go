package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupDBCRUD(r *gin.Engine) {
	// 查：获取所有学生
	r.GET("/students", func(c *gin.Context) {
		var students []Student
		// 查询所有记录，结果写入 students
		// 传地址，GORM 才能修改原变量
		DB.Find(&students)
		c.JSON(http.StatusOK, students)
	})

	// 查：获取单个学生
	r.GET("/students/:id", func(c *gin.Context) {
		var student Student

		// 按 id 查第一条记录
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
		// 把请求体 JSON 解析到 student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}
		//  插入一条新记录，自动填充 ID 和时间
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
		// GORM 默认是软删除，不是真正从数据库删除，而是把 DeletedAt 字段设为当前时间，之后查询会自动过滤掉这条记录。
		result := DB.Delete(&Student{}, c.Param("id"))
		// 实际影响的行数
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
	})
}

//
