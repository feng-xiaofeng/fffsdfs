// 中间件 middlewares包
package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	// 可以判断是否登录
	fmt.Println(time.Now())

	// 调用请求剩余的程序
	// c.Next()
	c.Set("name", "婷婷") //中间件对应控制器间共享数据
	// 定义一个goroutine统计日志
	cCp := c.Copy() //必须先Copy、
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}

func AitcMiddleware(c *gin.Context) {
	fmt.Println(time.Now())
	// 终止执行程序
	c.Abort()
	fmt.Println("time.Now()")
}
