package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/dd", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"ดดดดด": "พี",
		//})
		c.String(200, "success")
	})

	r.Run(":6001")
}

