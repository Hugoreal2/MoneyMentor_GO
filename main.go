package main

import (
	"Main/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up routes
	http.SetupRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
