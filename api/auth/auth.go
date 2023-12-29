package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"KlinikRidsu/api/auth/login"
	"KlinikRidsu/api/auth/register"
	"KlinikRidsu/api/auth/logout"
	"KlinikRidsu/api/auth/user"
)

func Auth(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
	register.Register(r, db)
	logout.Logout(r)
	user.User(r, db)
}
