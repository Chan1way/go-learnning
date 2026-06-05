package handler

import (
	"go-learnning/db"
	"go-learnning/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		Fail(c, 400, "参数错误")
		return
	}
	if db.DB.Create(&user).Error != nil {
		Fail(c, 500, "注册失败")
		return
	}
	Success(c, nil)
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		Fail(c, 400, "参数错误")
		return
	}

	// 查用户
	var user model.User
	if db.DB.Where("username = ? AND password =?",
		input.Username, input.Password).First(&user).Error != nil {
		Fail(c, 401, "用户名或密码错误")
		return
	}

	// 第二步：生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		Fail(c, 500, "生成token失败")
		return
	}

	// 第三步：返回 token
	Success(c, gin.H{"token": tokenString})
}
