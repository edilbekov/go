package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	ID uint
	Name string
	Image string
}

var err error
const dsn = "root:root@tcp(mysql:3306)/go"
var db *gorm.DB

func init(){
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Bazaga qo'shilmadi"))
	}

	if err := db.AutoMigrate(&Product{}); err != nil {
		panic(fmt.Sprintf("migrate bo'lmadi"))
	}
}

func GetProducts(c *gin.Context) {
	// var products []Product
	c.JSON(http.StatusOK, [])
}