package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}
