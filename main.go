package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"KlinikRidsu/databases"
	"KlinikRidsu/api"
	"KlinikRidsu/auth"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func Routes(r *gin.Engine, db *gorm.DB) {
	auth.SetupRoutes(r, db)
	api.Jadwal(r, db)
	api.Reservasi(r, db)
}

func main() {
	db := databases.InitDatabase()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://xhaeffer.me:11092",
		"http://localhost:3000",
	}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))
	r.Use(gin.Recovery())
	Routes(r, db)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
