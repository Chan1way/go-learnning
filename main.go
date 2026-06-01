package main

import (
	"go-learnning/db"
	"go-learnning/router"
)

func main() {
	db.Init()
	r := router.Setup()
	r.Run(":8080")
}
