package main

import (
	"github.com/gin-gonic/gin"
	route "images/router"
)

func main() {
	router := gin.Default()
	route.Init(router)
	router.Run(":8080")
}
