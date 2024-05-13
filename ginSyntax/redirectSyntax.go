package ginSyntax

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func redirectTest() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.Run()
}
