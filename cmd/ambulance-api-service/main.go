package main

import (
	"github.com/MatejLednicky/ambulance-webapi/api"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
    "github.com/MatejLednicky/ambulance-webapi/internal/ambulance_wl"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("AMBULANCE_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())

	// request routings
	ambulance_wl.AddRoutes(engine)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
