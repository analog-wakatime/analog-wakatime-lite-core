package main

import (
	"analog-wakatime-lite-core/config"
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func testrouters(app *gin.RouterGroup) {
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
	v1Docs := app.Group("/api/v1/docs")
	{
		v1Docs.GET("", func(c *gin.Context) {
			c.File("swagger/index.html")
		})

		v1Docs.GET("/", func(c *gin.Context) {
			c.File("swagger/index.html")
		})

		v1Docs.GET("/openapi.yaml", func(c *gin.Context) {
			c.File("api.yaml")
		})
	}

	v1 := app.Group("/api/v1")
	{
		testrouters(v1)

	}

	app.Use(cors.New(config.CorsConfig()))

	app.Run(config.ConfigEnv())
}
