package main

import (
	"analog-wakatime-lite-core/config"
	"analog-wakatime-lite-core/core/api/auth"
	"fmt"
	"io"
	"os"

	"analog-wakatime-lite-core/db"
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
	err := db.InitRedis()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
	if err := db.InitDB(config.ConfigGetDatabaseURL()); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
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
		v1.POST("/register", auth.Register)
		v1.POST("/login", auth.Login)
	}

	app.Use(cors.New(config.CorsConfig()))

	fmt.Println("Starting server on port " + config.GetAppPort())
	app.Run(config.GetAppPort())
}
