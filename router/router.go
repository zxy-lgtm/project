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
		g1.PUT("/students/:uid", handler.Students)
		g1.PUT("/patriarchs", handler.Patriarchs)

		//获取用户信息
		g1.GET("/students/:uid", handler.StudentInfo)
		g1.GET("/patriarchs", handler.PatriarchsInfo)
		g1.GET("/time/:uid", handler.TimeInfo)
		g1.GET("/subject/:uid", handler.SubjectInfo)
		g1.GET("/comment/:uid", handler.CommentInfo)
		g1.GET("/award/:uid", handler.AwardInfo)
		g1.GET("/teach/:uid", handler.TeachInfo)
	}
	//课堂
	g2 := r.Group("/project/v1/class")
	{
		//新建课堂
		g2.POST("/", handler.ClassPost)

		//取消课堂（可能有，没看到，先不写）
		//g2.PUT("/")

		//查询课堂
		g2.GET("/", handler.ClassInfo)

		//更新课堂信息
		g2.PUT("/", handler.ClassPut)

	}
	/*//花园
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
	r.GET("/project/v1/homework")*/
}
