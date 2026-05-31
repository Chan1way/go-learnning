package main

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // 用内存模拟数据库
// var studentList = []Student{
// 	{Name: "小明", Score: 95},
// 	{Name: "小红", Score: 88},
// }

// func setupCRUD(r *gin.Engine) {
// 	// 查：获取所有学生
// 	r.GET("/students", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, studentList)
// 	})

// 	//查：去获取单个学生
// 	r.GET("/students/:id", func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil || id < 0 || id >= len(studentList) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "id 不合法"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, studentList[id])
// 	})

// 	// 增：新增学生
// 	r.POST("/students", func(c *gin.Context) {
// 		var s Student
// 		if err := c.ShouldBindJSON(&s); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
// 			return
// 		}
// 		studentList = append(studentList, s)
// 		c.JSON(http.StatusOK, gin.H{"message": "新增成功", "student": s})
// 	})

// 	// 改：修改学生
// 	r.PUT("/students/:id", func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil || id < 0 || id >= len(studentList) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "id 不合法"})
// 			return
// 		}
// 		var s Student
// 		if err := c.ShouldBindJSON(&s); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
// 		}
// 		studentList[id] = s
// 		c.JSON(http.StatusOK, gin.H{"message": "修改成功", "student": s})
// 	})

// 	// 删：删除学生
// 	r.DELETE("/students/:id", func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil || id < 0 || id >= len(studentList) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "id不合法"})
// 			return
// 		}
// 		studentList = append(studentList[:id], studentList[id+1:]...)
// 		c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
// 	})
// }

// //   Gin 把请求和响应合并成一个 c *gin.Context：

// //   ┌─────────────────┬──────────────────────┐
// //   │      操作        │         方法          │
// //   ├─────────────────┼──────────────────────┤
// //   │ 取路径参数        │ c.Param("id")        │
// //   ├─────────────────┼──────────────────────┤
// //   │ 取 URL 参数      │ c.Query("name")      │
// //   ├─────────────────┼──────────────────────┤
// //   │ 解析请求体 JSON   │ c.ShouldBindJSON(&s) │
// //   ├─────────────────┼──────────────────────┤
// //   │ 返回文本          │ c.String(200, "...") │
// //   ├─────────────────┼──────────────────────┤
// //   │ 返回 JSON        │ c.JSON(200, data)    │
// //   └─────────────────┴──────────────────────┘

// //   ---
// //   四个 CRUD 接口

// //   ┌────────┬────────┬───────────────┐
// //   │  操作   │  方法   │     路径      │
// //   ├────────┼────────┼───────────────┤
// //   │ 查所有  │ GET    │ /students     │
// //   ├────────┼────────┼───────────────┤
// //   │ 查单个  │ GET    │ /students/:id │
// //   ├────────┼────────┼───────────────┤
// //   │ 新增    │ POST   │ /students     │
// //   ├────────┼────────┼───────────────┤
// //   │ 修改    │ PUT    │ /students/:id │
// //   ├────────┼────────┼───────────────┤
// //   │ 删除    │ DELETE │ /students/:id │
// //   └────────┴────────┴───────────────┘
