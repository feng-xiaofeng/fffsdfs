// 后台公用模块

package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseConterller struct{}

type IndexUserConfig struct {
	BaseConterller
}
type ArtocdeUserConfig struct {
	BaseConterller
}

func (con IndexUserConfig) Index(c *gin.Context) {
	con.error(c)
}
func (con BaseConterller) success(c *gin.Context) {
	c.String(http.StatusOK, "成功")
}
func (con BaseConterller) error(c *gin.Context) {
	c.String(http.StatusOK, "失败")
}
