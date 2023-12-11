package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"KlinikRidsu/server/session"
)

func CekReservasi (r *gin.Engine) {
	r.GET("/cekreservasi", session.VerifyToken(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "cekreservasi.html", nil)
	})
}