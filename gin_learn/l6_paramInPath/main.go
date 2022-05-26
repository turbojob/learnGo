package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   int    `uri:"id"`
	name string `uri:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/:id/:name", func(c *gin.Context) {
		//name := c.Param("name")
		//id := c.Param("id")
		var p Person
		if err := c.ShouldBindUri(&p); err != nil {
			c.Status(404)
			return
		}
		c.JSON(200, gin.H{
			"name": p.name,
			"id":   p.ID,
		})

	})
	r.Run()
}
