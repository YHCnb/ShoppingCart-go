/**
* @author:YHCnb
* Package:
* @date:2023/9/24 16:59
* Description:
 */
package database

import (
	"ShoppingCart/util/config"
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

// User 用户
type User struct {
	Sid      uint   `gorm:"primarykey" json:"sid"`
	Username string `gorm:"not null;uniqueIndex" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Level    uint   `gorm:"not null" json:"level"`
}

// GoodInfo 商品信息
type GoodInfo struct {
	Id    uint    `gorm:"primarykey" json:"id"`
	Name  string  `gorm:"not null" json:"name"`
	Price float64 `gorm:"not null" json:"price"`
	Intro string  `json:"intro"`
	Pic   string  `json:"pic"`
}

// GoodItem 商品记录
type GoodItem struct {
	ID  uint `json:"id"`
	Num int  `json:"number"`
}

// GoodInfoItem 详细商品记录
type GoodInfoItem struct {
	GoodInfo GoodInfo `json:"good"`
	Num      int      `json:"number"`
}

// Cart 购物车
type Cart struct {
	Sid   uint   `gorm:"primarykey" json:"sid"`
	Goods string `json:"goods"`
}

// Order 订单
type Order struct {
	Id    uint      `gorm:"primarykey" json:"id"`
	Sid   uint      `gorm:"nut null" json:"sid"`
	Goods string    `json:"goods"`
	Price float64   `gorm:"not null" json:"price"`
	Time  time.Time `gorm:"autoCreateTime" json:"time"`
}

// GoodsToJson GoodInfoItems转Json
func GoodsToJson(goods []GoodInfoItem) string {
	if len(goods) == 0 {
		return "[]"
	}

	var goodsItems []GoodItem
	for i := 0; i < len(goods); i++ {
		goodsItems = append(goodsItems, GoodItem{
			ID:  goods[i].GoodInfo.Id,
			Num: goods[i].Num,
		})
	}
	goodsJSON, err := json.Marshal(goodsItems)
	if err != nil {
		panic(err)
	}
	return string(goodsJSON)
}

// JsonToGoods Json转GoodInfoItems
func JsonToGoods(goodsJSON string) []GoodInfoItem {
	var goodsItems []GoodItem
	var goods []GoodInfoItem
	if goodsJSON == "[]" {
		return goods
	}
	if err := json.Unmarshal([]byte(goodsJSON), &goodsItems); err != nil {
		panic(err)
	}
	for i := 0; i < len(goodsItems); i++ {
		var good GoodInfo
		DB.First(&good, goodsItems[i].ID)
		goods = append(goods, GoodInfoItem{
			GoodInfo: good,
			Num:      goodsItems[i].Num,
		})
	}
	return goods
}

func Init() {
	dsn := config.Config.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	db.AutoMigrate(&User{}, &GoodInfo{}, &Cart{}, &Order{})
}

// Test 数据库测试
func Test() {

}
