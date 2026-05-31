package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 重新定义 Student，加上gorm需要的字段
type Student struct {
	gorm.Model        // 自动加 ID、CreatedAt、UpdatedAt、DeletedAt
	Name       string `json:"name"`
	Score      int    `json:"score"`
}

func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	// 自动建表
	DB.AutoMigrate(&Student{})
}
