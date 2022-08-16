package routers

import (
	"my_demo/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/")
	{
		adminRouters.GET("/", admin.Homepage)
		adminRouters.GET("/denglu", admin.DefaultDengLu)
		adminRouters.GET("/denglu/zhuce", admin.ApiZhuCe)
	}
}
