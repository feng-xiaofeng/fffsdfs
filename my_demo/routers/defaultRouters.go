package routers

import (
	"my_demo/controllers/inying"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.POST("/denglu/zhuce", inying.DefaultZhuCes)
	}
}
