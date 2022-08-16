package api

import (
	"github.com/gin-gonic/gin"
)

func ApiDengLu(c *gin.Context) {
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	// []models.User(Username:username, Password:password)
	// models.DB.Where("username=?", username).Find(&models.DB.)
	//校验用户名是否存在以及密码是否正确
	string sql := "select CC_AutoId,CC_UserName,CC_LoginPassword from T_Seats where cc_loginId=@uid";

}
