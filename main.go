package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mzumi/gin-sample/handler"
)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.POST("/article", handler.ArticleCreate())

	return r
}
