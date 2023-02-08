package main

import (
	"ginChat/routers"
	"ginChat/utils"
)

func main() {
	// 初始化读取配置文件
	utils.InitConfig()

	// 初始化mysql数据库连接
	utils.InitMySQL()

	// 初始化redis数据库
	utils.InitRedis()

	// 初始化路由
	routers.InitRouters()
}
