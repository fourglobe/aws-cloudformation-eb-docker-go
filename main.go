package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	/// !WARN: don't touch me.
	r.GET("/", func(c *gin.Context) {
		c.JSON(204, nil)
	})

	// mapping url.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
