package handler

import (
	"project/model"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	var tmpinfo RegisterInfo
	if err := c.BindJSON(&tmpinfo); err != nil {
		c.JSON(400, gin.H{
			"message": "输入格式有误",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "注册成功！",
	})
}

func Students(c *gin.Context) {
	var tmpstudent model.Students
	if err := c.BindJSON(&tmpstudent); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "修改成功！",
	})

}

func s(c *gin.Context) {
	var tmpstudent model.Students
	if err := c.BindJSON(&tmpstudent); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "修改成功！",
	})

}

type RegisterInfo struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Tel    string `json:"tel"`
	Email  string `json:"email"`
}
