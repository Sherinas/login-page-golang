package main

import (
	"log"
	"main/Handler"
	auth "main/middleware"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	db := models.DatabaseSetup()
	if db == nil {
		log.Fatalf("Failed to connect to database")
	}
	models.SetDatabase(db)
	Handler.SetDatabase(db)
	auth.SetDatabase(db)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Next()
	})

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**")

	router.GET("/", Handler.SignIN)
	router.POST("/", auth.LoginHandler)
	router.GET("/home", Handler.GetHome)
	router.GET("/register", Handler.RegForm)

	router.POST("/register", Handler.RegisterHandler)
	router.GET("/logout", Handler.LogOut)

	router.GET("/admin", Handler.AdminGet)

	router.GET("/create", Handler.UserCreateFrmAdmin)

	router.POST("/create", Handler.UserCreate)

	router.POST("/delete", Handler.UserDelete)

	router.GET("/search", Handler.UserSerch)

	router.POST("/edit", Handler.UserEdit)

	router.Run(":8000")
}
