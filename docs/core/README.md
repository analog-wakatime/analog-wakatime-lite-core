# At the moment, the Core has such a routes:

```
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

	auth := app.Group("/api/v1")
	auth.Use(auths.JWTAuthMiddleware())
	{
		auth.GET("/my-info", profiles.GetMyInfo)
	}


```

## Routers fro Docs:

go to 

```
http://localhost:8080/
```
to see the docs

```
- GET: api/v1/docs

```