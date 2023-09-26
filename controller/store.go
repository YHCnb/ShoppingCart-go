/**
* @author:YHCnb
* Package:
* @date:2023/9/25 11:06
* Description:
 */
package controller

import (
	"ShoppingCart/database"
	"ShoppingCart/util/config"
	"github.com/gin-gonic/gin"
)

// GoodListQuery 获取商品列表请求结构，要有分页
type GoodListQuery struct {
	Page  int    `form:"page"`
	Order string `form:"order"` //rand | new
}

// GoodList 获取商品列表
func GoodList(c *gin.Context) {
	var query GoodListQuery
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误awa"})
		return
	}
	var goods []database.GoodInfo
	q := database.DB.Model(&database.GoodInfo{}).Select("id,name,intro,price,pic")

	if query.Order == "rand" {
		q = q.Order("random()")
	} else { //默认new
		q = q.Order("id DESC")
	}

	page_size := int(config.Config.GoodPageSize)
	q = q.Limit(page_size).Offset((query.Page - 1) * page_size)
	q.Find(&goods)

	c.JSON(200, goods)
}
