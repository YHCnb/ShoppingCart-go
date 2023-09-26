/**
* @author:YHCnb
* Package:
* @date:2023/9/25 10:41
* Description:
 */
package middleware

import (
	"ShoppingCart/util/config"
	"ShoppingCart/util/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CheckLogin 验证用户是否登录
func CheckLogin(strict bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("fake-cookie")
		uid, ok, admin := jwt.VerifyUserToken(token, config.Config.Key)
		if ok {
			c.Set("uid", uid)
			uid_uint, err := strconv.ParseUint(uid, 10, 32)
			if err != nil {
				c.JSON(500, gin.H{"msg": "获取用户ID错误Orz"})
				c.Abort()
				return
			}
			c.Set("uid_uint", uint(uid_uint))
			c.Set("admin", admin)
		} else if strict {
			c.JSON(401, gin.H{"msg": "请先登录awa"})
			c.Abort()
		}
	}
}
