package ginSyntax

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func rspHtml() {
	r := gin.Default()
	r.LoadHTMLGlob("ginSyntax/html/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
	})
	r.Run()
}
