/**
* @author:YHCnb
* Package:
* @date:2023/9/25 11:03
* Description:
 */
package controller

import (
	"ShoppingCart/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetCart 获取购物车
func GetCart(c *gin.Context) {
	uid, _ := c.Get("uid_uint")
	var cart database.Cart
	fmt.Println("uid", uid.(uint))
	database.DB.Where("sid = ?", uid.(uint)).First(&cart)
	goods := database.JsonToGoods(cart.Goods)
	c.JSON(200, goods)
}

// AddGood 添加商品
func AddGood(c *gin.Context) {
	var queries []database.GoodItem
	if err := c.ShouldBind(&queries); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误awa"})
		return
	}

	uid, _ := c.Get("uid_uint")
	var cart database.Cart
	database.DB.Where("sid = ?", uid).First(&cart)
	goods := database.JsonToGoods(cart.Goods)
	// 遍历queries
	for i := 0; i < len(queries); i++ {
		var flag bool
		fmt.Println("len(goods)", len(goods))
		for j := 0; j < len(goods); j++ {
			if goods[j].GoodInfo.Id == queries[i].ID {
				goods[j].Num += queries[i].Num
				flag = true
				break
			}
		}
		fmt.Println("len(goods)", len(goods))
		if !flag {
			var good database.GoodInfo
			database.DB.First(&good, queries[i].ID)
			if good.Id == 0 {
				c.JSON(500, gin.H{"msg": "商品不存在Orz"})
				return
			}
			goods = append(goods, database.GoodInfoItem{GoodInfo: good, Num: queries[i].Num})
		}
	}
	cart.Goods = database.GoodsToJson(goods)
	database.DB.Save(&cart)
	c.JSON(200, goods)
}

// DeleteGood 删除商品
func DeleteGood(c *gin.Context) {
	var queries []database.GoodItem
	if err := c.ShouldBind(&queries); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误awa"})
		return
	}

	uid, _ := c.Get("uid_uint")
	var cart database.Cart
	database.DB.Where("sid = ?", uid).First(&cart)
	goods := database.JsonToGoods(cart.Goods)
	// 遍历queries
	for i := 0; i < len(queries); i++ {
		var flag bool
		for j := 0; j < len(goods); j++ {
			if goods[j].GoodInfo.Id == queries[i].ID {
				goods[j].Num -= queries[i].Num
				if goods[j].Num <= 0 {
					goods = append(goods[:j], goods[j+1:]...)
				}
				flag = true
				break
			}
		}
		if !flag {
			c.JSON(500, gin.H{"msg": "购物车中不存在此商品Orz"})
			return
		}
	}
	cart.Goods = database.GoodsToJson(goods)
	database.DB.Save(&cart)
	c.JSON(200, goods)
}

// ClearCart 清空购物车
func ClearCart(c *gin.Context) {
	uid, _ := c.Get("uid_uint")
	var cart database.Cart
	database.DB.Where("sid = ?", uid).First(&cart)
	cart.Goods = "[]"
	database.DB.Save(&cart)
	c.JSON(200, gin.H{"msg": "清空购物车成功OvO"})
}
