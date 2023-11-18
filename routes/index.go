package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index (r *gin.Engine, db *gorm.DB) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}