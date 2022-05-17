package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pang(c *gin.Context) {
	//// H is a shortcut for map[string]interface{}
	c.JSON(200, gin.H{
		"message": "pang",
	})
}
func main() {
	//Create an instance of Engine, by using New() or Default()
	//返回 Engine结构体
	//  pool             sync.Pool  存放context
	//	trees            methodTrees 存放路由处理函数
	r := gin.Default()

	r.GET("/ping", pang)
	r.GET("/hello", hello)
	r.GET("/testH", testH)
	//启动  默认端口8080  基于net http 包下的
	//修改端口
	r.Run(":1234")
}

func testH(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "自己的map",
	})
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
