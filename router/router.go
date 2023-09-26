/*
 * @Author: flwfdd
 * @Date: 2023-03-13 10:39:47
 * @LastEditTime: 2023-05-16 02:10:19
 * @Description: 路由配置
 */
package router

import (
	"ShoppingCart/controller"
	"ShoppingCart/middleware"
	"github.com/gin-gonic/gin"
)

// SetRouter 配置路由
func SetRouter(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "Welcome to ShoppingCart!"})
	})

	// 用户模块
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)

	// 商品模块
	router.GET("/goods", controller.GoodList)

	// 购物车模块
	router.POST("/cart", middleware.CheckLogin(true), controller.AddGood)
	router.GET("/cart", middleware.CheckLogin(true), controller.GetCart)
	router.DELETE("/cart", middleware.CheckLogin(true), controller.DeleteGood)

	// 订单模块
	router.POST("/order", middleware.CheckLogin(true), controller.DoOrder)
	router.GET("/orders", middleware.CheckLogin(true), controller.OrderList)
}
