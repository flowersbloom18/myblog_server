package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/comment/:blogID/:name", func(c *gin.Context) {
		id := c.Param("blogID")
		name := c.Param("name")
		fmt.Println("id=", id, name)
		// 在这里处理您的逻辑，使用获取到的 id 和 name 值

		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})

	router.Run(":8080")
}
