package routes

import (
	"github.com/edilbekov/go/app/controllersdaw/controllersdw"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")dawdaw

	users := api.Group("/users")
	{
		users.GET("/", controllerswdaw.GetUsers)
		users.POST("/", controllers.CreateUser)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	products := api.Group("/products")
	{
		products.GET("/", controllers.GetProducts)
		// products.POST("/", controllers.CreateProduct)
		// products.GET("/:id", controllers.GetProductByID)
		// products.PUT("/:id", controllers.UpdateProduct)
		// products.DELETE("/:id", controllers.DeleteProduct)
	}
}
