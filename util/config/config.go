/**
* @author:YHCnb
* Package:
* @date:2023/9/24 17:01
* Description:
 */
package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Dsn string
	Key string
	// 用于登录的token过期时间
	LoginExpire  int64 `yaml:"login_expire"`
	GoodPageSize uint  `yaml:"good_page_size"`
	ReleaseMode  bool  `yaml:"release_mode"`
	Port         string
	Saver        struct {
		MaxSize int64 `yaml:"max_size"`
	}
}{}

func Init() {
	path := "config.yml"
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("config.yml not found")
		os.Exit(1)
	}
	configor.Load(&Config, path)
}
