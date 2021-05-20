package router

import (
	"project/handler"

	"github.com/gin-gonic/gin"
)

//Router .
func Router(r *gin.Engine) {

	//user:
	g1 := r.Group("/project/v1/user")
	{
		//注册
		g1.POST("/in", handler.User)

		//登陆
		g1.POST("/up", handler.Userup)

		//修改用户信息
		g1.PUT("/students")
		g1.PUT("/patriarchs")

		//获取用户信息
		g1.GET("/students")
		g1.GET("/patriarchs")
	}
	//记事本
	g2 := r.Group("/project/v1/notepad")
	{
		//新建待办
		g2.POST("/create")

		//取消待办
		g2.PUT("/")

		//查询待办
		g2.GET("/")

		//消除未完成待办
		g2.PUT("/clear")

		//完成待办
		g2.POST("/")
	}
	//花园
	g3 := r.Group("/project/v1/garden")
	{
		//获取用户花园皮肤
		g3.GET("/")
		//g3.GET("/skin", handler.GetSkins)

		//新增皮肤
		g3.POST("/")
		//g3.POST("/buy", handler.BuySkin)

		//买花
		g3.PUT("/")
		//g3.PUT("/buy", handler.BuyFlower)
	}
	//作业
	r.GET("/project/v1/homework")
}
