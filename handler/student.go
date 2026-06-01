package handler

import (
	"go-learnning/db"
	"go-learnning/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询所有学生
func GetStudents(c *gin.Context) {
	var students []model.Student
	db.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

// 根据id查询学生
func GetStudent(c *gin.Context) {
	var student []model.Student
	result := db.DB.First(&student, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// 添加学生
func CreateStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
	}
	result := db.DB.Create(&student)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "新增成功", "student": student})
}

// 更改学生信息
func UpdateStudent(c *gin.Context) {
	var student []model.Student
	result := db.DB.First(*&student, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
	}
	db.DB.Save(&student)
	c.JSON(http.StatusOK, gin.H{"message": "修改成功", "student": student})
}

// 删除学生
func DeleteStudent(c *gin.Context) {
	result := db.DB.Delete(&model.Student{}, c.Param("id"))
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})

}
