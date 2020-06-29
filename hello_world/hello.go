package main

import "github.com/gin-gonic/gin"

func main() {

	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "hello go gin"})
	})

	engine.Run() // listen and serve on 0.0.0.0:8080
}
