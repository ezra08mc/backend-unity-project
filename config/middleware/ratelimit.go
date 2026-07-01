package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func GlobalRateLimiter(rps float64, burst int, skip map[string]struct{}) gin.HandlerFunc {
	if rps <= 0 {
		return func(c *gin.Context) { c.Next() }
	}
	if burst <= 0 { 
		burst = int(rps)
		if burst < 1 {
			burst = 1
		}
	}
	lim := rate.NewLimiter(rate.Limit(rps), burst)

	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		if skip != nil {
			if _, ok := skip[c.FullPath()]; ok {
				c.Next()
				return
			}
		}
		if !lim.Allow() {
			retryAfter := time.Until(time.Now().Add(lim.Reserve().Delay())) / time.Millisecond
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":          "rate limit exceeded",
				"retry_after_ms": retryAfter,
			})
			return
		}
		c.Next()
	}
}
