package http

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/list/:id", getAccount)
	router.POST("/accounts/:id/transaction", addTransaction)
}
