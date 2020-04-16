package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/torkashvand/goshortener/controllers"
	"github.com/torkashvand/goshortener/middlewares"
	"github.com/torkashvand/goshortener/models"
)

// SetupRouter provide router for application
func SetupRouter(dbProvider models.DatabaseProvide) *gin.Engine {
	router := gin.Default()

	// Provide db value for controller handlres
	db := dbProvider.GetDB()
	router.Use(middlewares.DBMiddlware(db))

	// Provide routing for application
	router.GET("/links", controllers.FindLinks)
	router.POST("/links", controllers.CreateLink)
	router.GET("/redirect/:shortcut", controllers.Redirect)

	return router
}
