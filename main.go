package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torkashvand/goshortener/cmd"
	"github.com/torkashvand/goshortener/controllers"
	"github.com/torkashvand/goshortener/models"
)

func main() {

	r := gin.Default()
	db := models.InitDB()
	defer db.Close()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/links", controllers.FindLinks)
	r.POST("/links", controllers.CreateLink)
	r.GET("/redirect/:shortcut", controllers.Redirect)

	r.Run()
	cmd.Execute()

}
