package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	config.loadEnv()

	r := gin.Default()

	r.Run(":8080")
}
