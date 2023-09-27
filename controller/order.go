/**
* @author:YHCnb
* Package:
* @date:2023/9/25 13:41
* Description:
 */
package controller

import (
	"ShoppingCart/database"
	"github.com/gin-gonic/gin"
	"time"
)

// DoOrder 下单
func DoOrder(c *gin.Context) {
	var queries []database.GoodItem
	if err := c.ShouldBind(&queries); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误awa"})
		return
	}
	uid, _ := c.Get("uid_uint")
	var cart database.Cart
	database.DB.Where("sid = ?", uid).First(&cart)
	cartGoods := database.JsonToGoods(cart.Goods)
	var orderGoods []database.GoodInfoItem
	var total float64

	// 遍历queries
	for i := 0; i < len(queries); i++ {
		var flag bool
		for j := 0; j < len(cartGoods); j++ {
			if cartGoods[j].GoodInfo.Id == queries[i].ID {
				orderGoods = append(orderGoods, database.GoodInfoItem{GoodInfo: cartGoods[j].GoodInfo, Num: queries[i].Num})
				total += cartGoods[j].GoodInfo.Price * float64(queries[i].Num)

				cartGoods[j].Num -= queries[i].Num
				if cartGoods[j].Num <= 0 {
					cartGoods = append(cartGoods[:j], cartGoods[j+1:]...)
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
	order := database.Order{
		Sid:   uid.(uint),
		Goods: database.GoodsToJson(orderGoods),
		Price: total,
	}
	cart.Goods = database.GoodsToJson(cartGoods)
	database.DB.Save(&cart)

	database.DB.Create(&order)
	database.DB.Where("sid = ?", uid).Last(&order)
	c.JSON(200, gin.H{"id": order.Id, "goods": orderGoods, "price": total, "time": order.Time})
}

type OrderResponse struct {
	ID    uint                    `json:"id"`
	Goods []database.GoodInfoItem `json:"goods"`
	Price float64                 `json:"price"`
	Time  time.Time               `json:"time"`
}

// OrderList 获取历史Order列表
func OrderList(c *gin.Context) {
	uid, _ := c.Get("uid_uint")
	var orders []database.Order
	database.DB.Where("sid = ?", uid).Find(&orders)

	var orderList []OrderResponse
	for i := 0; i < len(orders); i++ {
		goods := database.JsonToGoods(orders[i].Goods)
		orderList = append(orderList, OrderResponse{
			ID:    orders[i].Id,
			Goods: goods,
			Price: orders[i].Price,
			Time:  orders[i].Time,
		})
	}
	c.JSON(200, orderList)
}
