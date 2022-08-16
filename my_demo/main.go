package main

import (
	"github.com/gin-gonic/gin"
	"my_demo/routers"
)

func main() {
	r := gin.Default()

	// 加载模板
	r.LoadHTMLGlob("template/**/*")
	// 配置静态web目录， 第一个表示路由， 第二个表示目录
	r.Static("/static", "./static")
	// FET 请求
	routers.AdminRoutersInit(r)
	// 获取表单post数据
	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)
	r.Run(":80")
}
