package router

import (
	"hello-go/controller/graphql"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// GET方法用于支持GraphQL的web界面操作
	// 如果不需要web界面，可根据自己需求用GET/POST或者其他都可以
	Router.POST("/graphql", graphql.GraphqlHandler())
	Router.GET("/graphql", graphql.GraphqlHandler())
}
