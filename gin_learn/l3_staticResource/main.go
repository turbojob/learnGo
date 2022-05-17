package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载静态文件  比如图片
	//1 加载路径所有的  内部还是用的staticfs
	r.Static("/images", "./images")
	//2
	// StaticFS works just like `Static()`
	// but a custom `http.FileSystem` can be used instead.
	// Gin by default user: gin.Dir()
	//fs 是http.FileSystem 类型的
	r.StaticFS("/static", http.Dir("./static"))
	//3  加载单独的静态文件
	r.StaticFile("/index", "index.html")
	r.Run()
}
