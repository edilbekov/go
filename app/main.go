package main

import (
	"github.com/gin-gonic/gin"
	"github.com/edilbekov/go/app/routes"
	"github.com/davecgh/go-spew/spew"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	spew.Dump(1)
	// router := gin.Default()

	// routes.SetupRoutes(router)

	// router.Run(":8000")
}