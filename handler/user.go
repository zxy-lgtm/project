package handler

import (
	"fmt"
	"log"
	"project/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//测试数据
var Student = model.Students{
	ID:           123,
	Name:         "名字",
	IdentityCard: "身份证",
	School:       "学校",
	Major:        "专业",
	Tel:          "电话",
	Grade:        "年级",
	Gender:       1,
	Subject:      "专业",
	Type:         1,
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

//注册
func User(c *gin.Context) {
	//var tmpinfo RegisterInfo
	var tmpinfo model.LoginInfo
	if err := c.BindJSON(&tmpinfo); err != nil {
		c.JSON(400, gin.H{
			"message": "输入格式有误",
			"code":    "100001",
		})
		return
	}

	_, ok := model.FindUserAccount(tmpinfo.Account)
	if ok {
		c.JSON(200, gin.H{
			"message": "用户名重复！",
			"code":    "200001",
		})
		return
	}

	ok = model.CreateU(tmpinfo)
	if !ok {
		c.JSON(400, gin.H{
			"message": "存入数据失败！",
			"code":    "100002",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "注册成功！",
		"code":    "200000",
	})
}

//登录
func Userup(c *gin.Context) {
	var tmpinfo model.LoginInfo
	if err := c.BindJSON(&tmpinfo); err != nil {
		c.JSON(400, gin.H{
			"message": "输入格式有误",
			"code":    "100001",
		})
		return
	}

	loginfo, ok := model.FindUserAccount(tmpinfo.Account)

	if !ok {
		c.JSON(404, gin.H{
			"message": "账号错误！",
			"code":    "100003",
		})
		return
	}

	if tmpinfo.Password != loginfo.Password {
		c.JSON(404, gin.H{
			"message": "密码错误！",
			"code":    "100004",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "登录成功！",
		"code":    "200000",
		"user_id": loginfo.ID,
	})
}

//完善信息
func Students(c *gin.Context) {

	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	var tmpstudent model.ChangeStudents
	if err := c.BindJSON(&tmpstudent); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	student, ok := model.FindStudentsID(id)

	if !ok {
		c.JSON(404, gin.H{
			"message": "找不到该id的用户！",
			"code":    "100003",
		})
	}
	student = model.Students{
		School:        tmpstudent.School,
		IdentityCard:  tmpstudent.IdentityCard,
		IdentityCardF: tmpstudent.IdentityCardF,
		IdentityCardB: tmpstudent.IdentityCardB,
		Name:          tmpstudent.Name,
		Major:         tmpstudent.Major,
		Money:         tmpstudent.Money,
		Grade:         tmpstudent.Grade,
	}

	ok = model.UpdateStudent(student)

	if !ok {
		c.JSON(400, gin.H{
			"message": "存入数据失败！修改失败！",
			"code":    "100002",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "修改成功！",
		"code":    "200000",
	})

}

//修改信息
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

func PatriarchsInfo(c *gin.Context) {
	var tmppatriarchs model.Patriarchs
	if err := c.BindJSON(&tmppatriarchs); err != nil {
		c.JSON(400, gin.H{
			"message": "id格式错误！",
		})
		return
	}
	if tmppatriarchs.ID == Patriarch.ID {
		c.JSON(200, gin.H{
			"message": Patriarchs,
		})
		return
	} else {
		c.JSON(404, gin.H{
			"message": "未找到该用户信息",
		})
	}

}

func StudentInfo(c *gin.Context) {

	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	student, ok := model.FindStudentsID(id)

	if !ok {
		c.JSON(404, gin.H{
			"message": "找不到该id的用户！",
			"code":    "100003",
		})
	}

	c.JSON(200, gin.H{
		"message": student,
		"code":    "200000",
	})

}

//time是教学时长哦
func TimeInfo(c *gin.Context) {
	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	times, ok := model.FindTimeUID(id)

	if !ok {
		c.JSON(404, gin.H{
			"message": "找不到该id的用户的相关信息！",
			"code":    "100003",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": times,
		"code":    "200000",
	})
}

func SubjectInfo(c *gin.Context) {
	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	subjects, ok := model.FindSubjectUID(id)

	if !ok {
		c.JSON(404, gin.H{
			"message": "找不到该id的用户的相关信息！",
			"code":    "100003",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": subjects,
		"code":    "200000",
	})

}

func CommentInfo(c *gin.Context) {
	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	comment, ok := model.FindCommentUID(id)

	if !ok {
		c.JSON(404, gin.H{
			"message": "找不到该id的用户的相关信息！",
			"code":    "100003",
		})
		return
	}

	var comment_return = model.CommentReturn{
		Grade: model.StrName(comment.Name),
		W1:    comment.W1,
		W2:    comment.W2,
		W3:    comment.W3,
		W4:    comment.W4,
		W5:    comment.W5,
		W6:    comment.W6,
		W7:    comment.W7,
		All:   comment.All,
		C1:    comment.C1,
		C2:    comment.C2,
		C3:    comment.C3,
		C4:    comment.C4,
		C5:    comment.C5,
		C6:    comment.C6,
		C7:    comment.C7,
	}

	c.JSON(200, gin.H{
		"message": comment_return,
		"code":    "200000",
	})

}

func AwardInfo(c *gin.Context) {
	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	awards, ok := model.FindAwardUID(id)

	if !ok {
		c.JSON(404, gin.H{
			"message": "找不到该id的用户的相关信息！",
			"code":    "100003",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": awards,
		"code":    "200000",
	})

}

func TeachInfo(c *gin.Context) {
	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
			"code":    "100001",
		})
		return
	}

	teaches, ok := model.FindTeachUID(id)

	if !ok {
		c.JSON(404, gin.H{
			"message": "找不到该id的用户的相关信息！",
			"code":    "100003",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": teaches,
		"code":    "200000",
	})

}

/*type RegisterInfo struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Tel    string `json:"tel"`
	Email  string `json:"email"`
}*/

func produceToken(uid string) string {
	//id, _ := strconv.Atoi(uid)
	claims := &model.JwtClaims{
		UID: uid,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	//fmt.Println(singedToken, err)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims model.JwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

var (
	key        = "miniProject" //salt
	ExpireTime = 3600          //token expire time
)
