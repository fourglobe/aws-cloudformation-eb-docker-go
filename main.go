package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/// !WARN: don't touch me.
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "ok",
		})
	})
	r.Run()
}
