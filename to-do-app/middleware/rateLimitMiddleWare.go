package middleware

import (
	"context"
	"log"
	"net/http"
	"strconv"

	config "to-do-app/config"

	"time"

	"github.com/gin-gonic/gin"
)

const (
	REDIS_KEY_FORMAT = "ratelimit:user:"
	RATELIMIT = 5
	RATELIMITTTL = time.Minute * 5
)


func RateLimitMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		getUserId, _ := c.Get("userId")
		user_id, _ := getUserId.(int)
		redisKey := REDIS_KEY_FORMAT + strconv.Itoa(user_id)

		exists, err := config.RedisClient.SetNX(ctx, redisKey, 1, RATELIMITTTL).Result()
		if err != nil {
			log.Println("Error getting request count from redis: ", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if !exists {
			log.Println("redis key exists", exists)
			reqCount, err := config.RedisClient.Get(ctx, redisKey).Int()
			if err != nil {
				log.Println("Error getting request count from redis: ", err)
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				c.Abort()
				return
			}

			config.RedisClient.Incr(ctx, redisKey)

			if reqCount >= RATELIMIT {
				log.Println("Rate limit exceeded")
				c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}