package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("welcome", welcomeHandler)
	r.GET("array", arrayHandler)
	r.GET("map", mapHandler)
	r.Run()
}

func mapHandler(c *gin.Context) {
	key := c.QueryMap("user")
	c.JSON(200, gin.H{
		"user=": key,
	})
}

func arrayHandler(c *gin.Context) {
	ids := c.QueryArray("ids")
	c.JSON(200, gin.H{
		"ids": ids,
	})
}

func welcomeHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "moren")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})

}
