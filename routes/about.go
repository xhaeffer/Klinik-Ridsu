package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func About (r *gin.Engine, db *gorm.DB) {
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
}