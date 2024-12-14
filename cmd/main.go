package main

import (
	"api-rate-limiter/internal/ratelimiter"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.Use(ratelimiter.RateLimiterMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
