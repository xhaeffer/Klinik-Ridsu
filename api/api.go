package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"KlinikRidsu/api/auth"
	"KlinikRidsu/api/jadwal"
	"KlinikRidsu/api/reservasi"
)

func API(r *gin.Engine, db *gorm.DB) {
	auth.Auth(r, db)
	jadwal.Jadwal(r, db)
	reservasi.Reservasi(r, db)
}
