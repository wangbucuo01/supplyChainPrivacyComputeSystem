package service

import (
	"strconv"
	"supplyChainPrivacyComputeSystem/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := model.User{}
	var err error
	// 用户注册需要提供用户名、密码、确认密码、身份、企业信息描述
	user.UserName = c.Request.FormValue("username")
	passwd := c.Request.FormValue("passwd")
	confirmpasswd := c.Request.FormValue("confirmpasswd")
	if user.UserName == "" || passwd == "" || confirmpasswd == "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户名或密码不能为空!",
			"data":    user,
		})
		return
	}

	if passwd != confirmpasswd {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "两次密码不一致!",
			"data":    user,
		})
		return
	}

	data := model.FindUserByName(user.UserName)
	if data.UserName != "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户名已注册",
			"data":    user,
		})
		return
	}

	user.Passwd = passwd
	// 身份：前端做 普通用户->2 的转换，这里传过来就是int
	user.Identity, err = strconv.Atoi(c.Request.FormValue("identity"))
	if err != nil {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "身份选择有误",
			"data":    user,
		})
		return
	}
	// 企业信息
	user.UserDesc = c.Request.FormValue("desc")
	user.State = 0
	model.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功!",
		"data":    user,
	})
}