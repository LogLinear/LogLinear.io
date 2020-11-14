package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	port string = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

func statusHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": true, "message": "Main app service works"})
}

func emailConnectionStatusHandler(c *gin.Context) {
	// Try to connect to email service
	resp, err := http.Get(fmt.Sprintf("http://emailservice%s/api/v1/testEmailService", port))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	if resp.StatusCode != 200 {
		c.AbortWithStatusJSON(500, gin.H{"error": "Non-200 Status code returned from email service"})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}

	var v interface{}
	json.Unmarshal(body, &v)
	data := v.(map[string]interface{})
	status := data["status"].(bool)
	message := data["message"].(string)
	if status != true && message != "Email service works" {
		c.AbortWithStatusJSON(500, gin.H{"error": "Main service failed to establish a connection with the email service"})
	}
	c.JSON(200, gin.H{"status": true, "message": "Main service can call email service"})
}
