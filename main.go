package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"projek1/databases"
	"projek1/routes"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func Routes(r *gin.Engine, db *gorm.DB) {
	routes.Auth(r)
	routes.Index(r)
	routes.About(r)
	routes.Jadwal(r, db)
	routes.Reservasi(r, db)
}

func main() {
	db := databases.InitDatabase()
	r := gin.Default()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/templates/css", "./templates/css")
	r.Static("/templates/scripts", "./templates/scripts")
	r.Static("/templates/img", "./templates/img")
	Routes(r, db)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}	
