package router

import "github.com/gin-gonic/gin"

func Init(route *gin.Engine) {
	upload(route)
	images(route)
}
