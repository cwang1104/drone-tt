package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/get", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	})

	r.Run(":56001")

}
