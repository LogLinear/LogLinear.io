package routes

import "github.com/gin-gonic/gin"

func testEmailServiceHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": true, "message": "Email service works"})
}
