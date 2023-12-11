package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func About (r *gin.Engine) {
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
}