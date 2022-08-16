package models

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	//实现跨域访问
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站，多个使用，分开，限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		//准许请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享， 确定共享
		AllowCredentials: true,
		//允许跨域的原点网站，可以直接return true 就万事大吉啦
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})
	e.Use(mwCORS)
}
