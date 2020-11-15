package routes

import "github.com/gin-gonic/gin"

//Register routes to the router group
func Register(router *gin.RouterGroup) {
	router.GET("/testEmailService", testEmailServiceHandler)
	router.GET("/sendSimpleMessage", simpleMessageHandler)
	router.POST("/sendVerification", sendVerificationHandler)
}
