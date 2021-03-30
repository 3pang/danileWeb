package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main1() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "alive",
		})
	})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run()
}
