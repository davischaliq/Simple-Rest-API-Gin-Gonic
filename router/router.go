package router

import (
	"telkomGin/handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", handler.Welcome)
	router.GET("/show", handler.ShowAllUser)
	// router.POST("/add", handler.AddUser)
	router.GET("/single/:id", handler.SingleUser)
	router.DELETE("/drop/:id", handler.DeleteUser)
}
