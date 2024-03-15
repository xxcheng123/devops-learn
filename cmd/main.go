package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/*name", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Hello World!", "timestamp": time.Now().UnixMilli()})
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
