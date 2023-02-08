package routers

import (
	"ginChat/controller"
	"ginChat/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouters() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/user/getuserlist", controller.GetUserList)
	r.POST("/user/registry", controller.CreateUser)
	r.POST("/user/login", controller.Login)
	r.POST("/user/delectuser", controller.DelectUser)
	r.POST("/user/updateuser", controller.UpdateUser)
	r.GET("/user/sendmsg", controller.SendMsg)
	r.GET("/user/sendusermsg", controller.SendUserMsg)
	r.POST("/user/searchfriends", controller.SearchFriends)
	r.POST("/attach/upload", controller.Upload)
	r.Run()
}
