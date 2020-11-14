package main

import (
	"emailservice/routes"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	port string = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

//SetupRouter sets up all the required router
func SetupRouter() *gin.Engine {
	app := gin.Default()

	// Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://dashboard.loglinear.io"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	apiGroup := app.Group("/api/v1")
	routes.Register(apiGroup)
	return app
}

func main() {
	app := SetupRouter()
	if port == ":" {
		port = ":8888"
	}
	app.Run(port)
}
