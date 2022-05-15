package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("Test")
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test Module  Gin",
		})
	})
	r.Run()
}
