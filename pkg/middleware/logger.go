package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		latency := time.Since(time.Now())
		log.Print(latency)

		status := c.Writer.Status()
		log.Print(status)
	}
}
