/**
* @author:YHCnb
* Package:
* @date:2023/9/24 17:47
* Description:
 */
package main

import (
	"ShoppingCart/database"
	"ShoppingCart/router"
	"ShoppingCart/util/config"
	"fmt"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	config.Init()
	database.Init()
	//database.Test()

	if config.Config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	app.Use(limits.RequestSizeLimiter(config.Config.Saver.MaxSize << 20))
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type", "fake-cookie", "webvpn-cookie"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		MaxAge:       12 * time.Hour,
	}))
	router.SetRouter(app)
	fmt.Println("BIT101-GO will run on port " + config.Config.Port)
	app.Run(":" + config.Config.Port)
}
