package ratelimiter

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	for {
		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			log.Println("Could not connect to Redis:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		log.Println("Connected to Redis!")
		break
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := fmt.Sprintf("rate_limiter:%s", clientIP)

		count, err := redisClient.Incr(context.Background(), key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			c.Abort()
			return
		}

		if count == 1 {
			redisClient.Expire(context.Background(), key, time.Minute)
		}

		if count > 10 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
