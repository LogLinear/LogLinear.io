package routes

import (
	"emailservice/email"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func testEmailServiceHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": true, "message": "Email service works"})
}

func simpleMessageHandler(c *gin.Context) {
	var v struct {
		Email string `json:"Email"`
	}
	c.ShouldBindBodyWith(&v, binding.JSON)
	recipient := v.Email
	if recipient == "" {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get email from payload"})
		return
	}
	id, err := email.SendSimpleMessage(recipient)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	log.Printf("Sent simplle message to %s", recipient)
	c.JSON(200, gin.H{"email id": id})
	return
}

func sendVerificationHandler(c *gin.Context) {
	var v struct {
		Email             string `json:"Email"`
		VerificationToken string `json:"VerificationToken"`
	}
	c.ShouldBindBodyWith(&v, binding.JSON)
	to := v.Email
	verificationToken := v.VerificationToken
	if to == "" && verificationToken == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email and verification token cannot be empty"})
		return
	}
	// Trigger the API call to email sending service
	resp, id, err := email.SendEmail(to, verificationToken)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Sending verification email to: %s", to)
	c.JSON(200, gin.H{"status": true, "message": id, "resp": resp})
}
