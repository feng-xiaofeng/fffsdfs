// 后台控制器 admin 包
package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
func DefaultDengLu(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/denglu.html", gin.H{})
}

// 注册
func ApiZhuCe(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/zhuce.html", gin.H{})
}
