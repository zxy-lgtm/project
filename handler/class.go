package handler

import (
	"project/model"

	"github.com/gin-gonic/gin"
)

//测试信息
var class0 = model.ClassInfo{
	ID:           1,
	Name:         "鸡兔同笼",
	StudentsName: "王芳",
	Status:       0,
	Tips:         "小学数学",
	Tips2:        "",
	Money:        0,
}

var class1 = model.ClassInfo{
	ID:           2,
	Name:         "呐喊赏析",
	StudentsName: "黄芳",
	Status:       1,
	Tips:         "初中语文",
	Tips2:        "七节课，45分钟/节",
	Money:        200,
}

//获取课堂信息
func ClassInfo(c *gin.Context) {
	var tmpclass model.ClassInfo
	if err := c.BindJSON(&tmpclass); err != nil {
		c.JSON(400, gin.H{
			"message": "id格式错误！",
		})
		return
	}
	if tmpclass.ID == Patriarch.ID {
		c.JSON(200, gin.H{
			"message": ClassInfo,
		})
		return
	} else {
		c.JSON(404, gin.H{
			"message": "未找到该课堂信息",
		})
	}
}

func ClassPut(c *gin.Context) {
	var tmpclass model.ClassInfo
	if err := c.BindJSON(&tmpclass); err != nil {
		c.JSON(400, gin.H{
			"message": "格式错误！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "修改成功！",
	})
}

func ClassPost(c *gin.Context) {
	var tmpclass model.ClassInfo
	if err := c.BindJSON(&tmpclass); err != nil {
		c.JSON(400, gin.H{
			"message": "格式错误！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "上传成功！",
	})
}
