package inying

import (
	"my_demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func DefaultZhuCes(c *gin.Context) {
	// var password1 string
	username := c.PostForm("username")
	password := c.PostForm("password1")
	password1 := c.PostForm("password2")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	if password != password1 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"crr": "两次输入的密码不一致",
		})
		return
	}
	//如果用户名没有传， 给一个10位数的随机字符串
	if len(username) == 0 {
		username = models.RandomString(10)
	}
	//数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号码必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}

	//判断手机是否存在
	// if models.CheckMobile(phone) {
	// c.JSON(http.StatusUnprocessableEntity, gin.H{
	// "code": 422,
	// "msg":  "用户已存在",
	// })
	// return
	// }
	models.VerifyEmailFormat(email)
	user := models.User{
		Username: username,
		Password: models.FromPassword(password),
		Phone:    phone,
		Email:    email,
	}
	models.DB.Create(&user)
	c.String(http.StatusOK, "注册成功")
}
