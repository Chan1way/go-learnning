package handler

import (
	"go-learnning/db"
	"go-learnning/model"

	"github.com/gin-gonic/gin"
)

// 查询所有学生
func GetStudents(c *gin.Context) {
	var students []model.Student
	db.DB.Find(&students)
	Success(c, students)
}

// 根据id查询学生
func GetStudent(c *gin.Context) {
	var student model.Student
	result := db.DB.First(&student, c.Param("id"))
	if result.Error != nil {
		Fail(c, 404, "学生不存在")
		return
	}
	Success(c, student)
}

// 添加学生
func CreateStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		Fail(c, 400, "参数错误")
		return
	}
	result := db.DB.Create(&student)
	if result.Error != nil {
		Fail(c, 500, "新增失败")
		return
	}
	Success(c, student)
}

// 更改学生信息
func UpdateStudent(c *gin.Context) {
	var student model.Student
	result := db.DB.First(&student, c.Param("id"))
	if result.Error != nil {
		Fail(c, 404, "学生不存在")
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		Fail(c, 400, "参数错误")
		return
	}
	db.DB.Save(&student)
	Success(c, student)
}

// 删除学生
func DeleteStudent(c *gin.Context) {
	result := db.DB.Delete(&model.Student{}, c.Param("id"))
	if result.RowsAffected == 0 {
		Fail(c, 404, "学生不存在")
		return
	}
	Success(c, nil)

}
