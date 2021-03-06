package website

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func websiteHandler(c *gin.Context) {
	isProduction := os.Getenv("is_production")
	log.Println(isProduction)
	// let react handle the website
	if isProduction == "false" {
		c.File("./website/build/index.html")
	} else {
		c.File("./build/index.html")
	}
}
