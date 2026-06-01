package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 声明一个全局变量 DB，整个程序都能用它操作数据库。
var DB *gorm.DB

// 重新定义 Student，加上gorm需要的字段
type Student struct {
	// gorm.Model 是 GORM 内置的结构体，嵌入后自动给你加 4 个字段：
	// ID        uint           // 主键，自动递增
	// CreatedAt time.Time      // 创建时间
	// UpdatedAt time.Time      // 更新时间
	// DeletedAt gorm.DeletedAt // 删除时间（软删除）
	// `json:"name"` 是结构体标签，告诉 JSON 序列化时用小写的 name，而不是 Name：
	// 没有标签时
	// {"Name":"小明","Score":95}
	// 有标签时
	// {"name":"小明","score":95}
	gorm.Model        // 自动加 ID、CreatedAt、UpdatedAt、DeletedAt
	Name       string `json:"name"`
	Score      int    `json:"score"`
}

func initDB() {
	var err error
	// 连接数据库
	// ┌────────────────────────────┬────────────────────────────────────────┐
	// │            部分             │                  意思                   │
	// ├────────────────────────────┼────────────────────────────────────────┤
	// │ gorm.Open                  │ 打开数据库连接                            │
	// ├────────────────────────────┼────────────────────────────────────────┤
	// │ sqlite.Open("students.db") │ 用 SQLite，数据存在 students.db 文件里    │
	// ├────────────────────────────┼────────────────────────────────────────┤
	// │ &gorm.Config{}             │ GORM 配置（用默认配置就够了）              │
	// └────────────────────────────┴────────────────────────────────────────┘
	DB, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		// panic 是 Go 里让程序立刻崩溃并报错的方式
		// 数据库连不上，程序没法继续跑，所以用 panic 直接停掉，而不是返回 error。
		panic("数据库连接失败：" + err.Error())
	}

	// 自动建表
	// 根据 Student struct 的字段，自动在数据库里创建对应的表。
	// 如果表已经存在，只会新增缺少的字段，不会删除已有数据。
	DB.AutoMigrate(&Student{})
}
