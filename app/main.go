package main

import (
	// _ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"github.com/gin-gonic/gin"
	// "routes"
	"github.com/edilbekov/go/app/routes"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	router := gin.Default()

	// API yo'llarini sozlash uchun alohida funksiyani chaqiramiz
	routes.SetupRoutes(router)

	router.Run(":8080")
}