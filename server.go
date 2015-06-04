package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hi":       "there",
			"integer":  123,
			"float":    1.23,
			"datetime": time.Now().Format(time.RFC3339Nano),
			"nested": gin.H{
				"I'm nested unicode!": "þæöóíí",
			},
		})
	})
	g.Run(":8080")
}
