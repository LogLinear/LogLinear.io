package service

import "github.com/gin-gonic/gin"

// Register all routes to this router group
func Register(router *gin.RouterGroup) {
	router.GET("/status", statusHandler)
	router.GET("/email", emailConnectionStatusHandler)
}
