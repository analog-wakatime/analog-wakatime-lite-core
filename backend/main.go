package main

import (
	"analog-wakatime-lite-core/config"
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func testrouters(app *gin.Engine) {
	app.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is working"})
	})
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is working"})
	})
}

func main() {
	gin.DisableConsoleColor()
	os.MkdirAll("logs", 0755)
	f, _ := os.OpenFile("logs/main_core.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f)
	app := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	testrouters(app)

	app.Use(cors.New(config.CorsConfig()))

	app.Run(config.ConfigEnv())
}
