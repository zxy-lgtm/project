package handler

import (
	"project/model"

	"github.com/gin-gonic/gin"
)

//测试数据
var Student = model.Students{
	ID:             123,
	Name:           "名字",
	IdentityCard:   "身份证",
	School:         "学校",
	Major:          "专业",
	Tel:            "电话",
	Grade:          "年级",
	Gender:         1,
	Subject:        "专业",
	AwardSituation: "获奖情况",
	Type:           1,
}

//测试数据
var Patriarch = model.Patriarchs{
	ID:                 123,
	NameParents:        "家长名字",
	NameChild:          "孩子名字",
	IdentityCardParent: "家长身份证",
	IdentityCardChild:  "孩子身份证",
	School:             "学校",
	Tel:                "电话",
	Grade:              "年级",
	GenderParents:      1,
	ScoreSituation:     "成绩情况",
}
var loginfo = Loginfo{
	ID:       "123",
	Password: "123",
}

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

func Userup(c *gin.Context) {
	var tmpinfo Loginfo
	if err := c.BindJSON(&tmpinfo); err != nil {
		c.JSON(400, gin.H{
			"message": "输入格式有误",
		})
		return
	}

	if tmpinfo.ID != loginfo.ID {
		c.JSON(404, gin.H{
			"message": "账号错误！",
		})
		return
	}

	if tmpinfo.Password != loginfo.Password {
		c.JSON(404, gin.H{
			"message": "密码错误！",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "登录成功！",
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

func Patriarchs(c *gin.Context) {
	var tmpstudent model.Patriarchs
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

type Loginfo struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
