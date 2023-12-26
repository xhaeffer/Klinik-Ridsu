package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"KlinikRidsu/auth/login"
	"KlinikRidsu/auth/register"
	"KlinikRidsu/auth/logout"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
	register.Register(r, db)
	logout.Logout(r)
}
