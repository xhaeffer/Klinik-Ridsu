package logout

import (
	"KlinikRidsu/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func logoutHandler(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	if _, ok := session.GetSession(c.Request)["user"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
		return
	}

	session.ClearSession(w, r)
	session.ClearToken(c)

	c.JSON(http.StatusOK, gin.H{ "message": "Logout Berhasil!",})
}

func Logout(r *gin.Engine) {
	r.GET("/logout", func(c *gin.Context) {
		logoutHandler(c.Writer, c.Request, c)
	})
}
