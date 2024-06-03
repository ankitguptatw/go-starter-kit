package securityheaders

import (
	"github.com/gin-gonic/gin"
)

func Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		ctx.Header("Cache-Control", "no-store")
		ctx.Header("Content-Security-Policy", "frame-ancestors 'none'")
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("X-XSS-Protection", "1; mode=block")

	}
}
