package db

import (
	"go-learnning/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	DB.AutoMigrate(&model.Student{})
}
