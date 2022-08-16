package routers

import (
	"my_demo/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/")
	{
		apiRouters.POST("/denglu", api.ApiDengLu)
	}
}
