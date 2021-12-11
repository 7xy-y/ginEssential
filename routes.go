package main

import (
	"github.com/gin-gonic/gin"
	"xietong.me/ginessential/controller"
	"xietong.me/ginessential/middleware"
	//"ginEssential/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/signup", controller.Signup)
	r.POST("/api/auth/upload", controller.UpLoad)
	//r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/auth/test", controller.Test)
	r.GET("/api/auth/info", controller.Info)
	r.POST("/api/auth/myinfo", controller.Myinfo)
	r.POST("/api/auth/send", controller.Send)
	return r
}
