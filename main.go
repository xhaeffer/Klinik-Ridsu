package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"KlinikRidsu/databases"
	"KlinikRidsu/routes"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func Routes(r *gin.Engine, db *gorm.DB) {
	routes.Login(r, db)
	// routes.Register(r, db)
	// routes.Index(r)
	// routes.About(r)
	routes.Jadwal(r, db)
	routes.Reservasi(r, db)
	// routes.CekReservasi(r)
}

func main() {
	db := databases.InitDatabase()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Update with your React app's URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	r.Use(gin.Recovery())
	// r.LoadHTMLGlob("templates/*.html")
	// r.Static("/templates/css", "./templates/css")
	// r.Static("/templates/scripts", "./templates/scripts")
	// r.Static("/templates/img", "./templates/img")
	Routes(r, db)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
