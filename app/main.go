package main

import (
	"github.com/gin-gonic/gin"
	"github.com/edilbekov/go/app/routes"
)

func main() {
	router := gin.Default()	

	routes.SetupRoutes(router)

	router.Run("0.0.0.0:8080")
}