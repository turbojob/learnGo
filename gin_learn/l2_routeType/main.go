package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		c.String(200, "GET")
	})
	r.POST("/posts", func(c *gin.Context) {
		c.String(200, "POST")
	})
	r.PUT("/posts/:id", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("PUT id:%s", c.Param("id")))
	})
	r.DELETE("/posts/:id", func(c *gin.Context) {
		c.String(200, "DELETE")
	})

	//匹配所有
	r.Any("/users", func(c *gin.Context) {
		c.String(200, "any")
	})
	r.Run()
}
