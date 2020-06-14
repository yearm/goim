package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Accept", "Accept-Encoding", "Accept-Language", "Connection", "Host", "Origin", "Referer", "Content-Type", "Content-Length", "User-Agent", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{},
		MaxAge:           12 * time.Hour,
	})
}

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		// TODO Authorization
	}
}
