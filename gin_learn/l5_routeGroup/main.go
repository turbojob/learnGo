package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//r.GET("/posts", GetHandler)
	//r.POST("/posts", PostHandler)
	////删除id = 1的
	//r.DELETE("/posts/1", DeleteHandler)

	//改造
	p := r.Group("/posts")
	{
		p.GET("/", GetHandler)
		p.POST("/", PostHandler)
		//删除id = 1的
		p.DELETE("/1", DeleteHandler)
	}

	r.Run()
}

func DeleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE",
	})

}

func PostHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "POST",
	})
}

func GetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET",
	})
}
