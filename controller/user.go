/**
* @author:YHCnb
* Package:
* @date:2023/9/25 10:35
* Description:
 */
package controller

import (
	"ShoppingCart/database"
	"ShoppingCart/util/config"
	"ShoppingCart/util/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login 登录
func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}
	user := database.User{}
	database.DB.Where("username = ?", loginRequest.Username).First(&user)
	if user.Sid == 0 {
		c.JSON(400, gin.H{"msg": "无此用户"})
		return
	}
	if user.Password == loginRequest.Password {
		token := jwt.GetUserToken(fmt.Sprint(user.Sid), config.Config.LoginExpire, config.Config.Key, user.Level == 0)
		fmt.Println("token", token)
		c.JSON(200, gin.H{"fake_cookie": token})
	} else {
		c.JSON(401, gin.H{"msg": "用户名或密码错误"})
	}
}

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Register 注册
func Register(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}
	user := database.User{}
	database.DB.Where("username = ?", registerRequest.UserName).First(&user)
	if user.Sid == 0 {
		database.DB.Create(&database.User{Username: registerRequest.UserName, Password: registerRequest.Password, Level: 1})
		// 插入空购物车
		database.DB.Create(&database.Cart{Sid: user.Sid, Goods: "[]"})
		c.JSON(200, gin.H{"msg": "注册成功OvO"})
	} else {
		c.JSON(401, gin.H{"msg": "用户名已存在"})
	}
}
