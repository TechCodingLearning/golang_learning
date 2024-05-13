package ginSyntax

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello www.topgoer.com!"})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}
