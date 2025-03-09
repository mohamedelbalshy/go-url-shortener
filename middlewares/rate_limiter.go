package middlewares

import (
	"context"
	"fmt"
	"go-url-shortener/database"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiterMiddleware limits requests per IP
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get client IP address
		clientIP := c.ClientIP()
		key := "rate_limit:" + clientIP

		// Define limits (e.g., 10 requests per minute)
		limit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT"))
		window, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_WINDOW")) // in seconds

		// Check request count in Redis
		count, err := database.RedisClient.Get(context.Background(), key).Int()
		if err != nil && err.Error() != "redis: nil" {
			fmt.Println("Error fetching rate limit:", err)
		}

		if count >= limit {
			// If exceeded, return 429 Too Many Requests
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests, please try again later."})
			c.Abort()
			return
		}

		// Increment request count
		err = database.RedisClient.Incr(context.Background(), key).Err()
		if err != nil {
			fmt.Println("Error incrementing rate limit:", err)
		}

		// Set expiration if it's a new counter
		if count == 0 {
			database.RedisClient.Expire(context.Background(), key, time.Duration(window)*time.Second)
		}

		// Allow request to proceed
		c.Next()
	}
}
