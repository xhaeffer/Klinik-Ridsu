package main

import (
	"KlinikRidsu/databases"
	"KlinikRidsu/api"
	"KlinikRidsu/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)

    configs.RecaptchaConfig()
	configs.JWTConfig()
	configs.SessionConfig()
}

func main() {
	db := databases.InitDatabase("local")
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

	api.API(r, db)
	r.Run()
}
