package controllers

import (
	"github.com/gin-gonic/gin"
)

func RouteConfig(c *gin.Context) {
	var routes []string
	routes = append(routes, c.Request.Method)
	routes = append(routes, c.Request.URL.Path)
	c.IndentedJSON(200, gin.H{
		"routes": routes,
	})

}
