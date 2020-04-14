package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/torkashvand/goshortener/cache"
	"github.com/torkashvand/goshortener/config"
	"github.com/torkashvand/goshortener/helpers"
	"github.com/torkashvand/goshortener/models"
)

// CreateLinkInput validate input data
type CreateLinkInput struct {
	Address string `json:"address" binding:"required"`
}

// FindLinks get all links
// GET /links
func FindLinks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var links []models.Link
	db.Find(&links)

	cfg := config.Config()
	for i := range links {
		baseAddress := cfg.GetString("BASE_REDIRECT_SERVER")
		links[i].Shortcut = baseAddress + links[i].Shortcut
	}
	c.JSON(http.StatusOK, gin.H{"data": links})
}

// CreateLink POST /links
// Create new links
func CreateLink(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateLinkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create link
	link := models.Link{Address: input.Address}
	db.Create(&link)
	link.Shortcut = helpers.ConvertBase(link.ID)
	db.Model(&link).Updates(link)
	cache.InitializeRedis()
	config := config.Config()
	cache.SetValue(link.Shortcut, link.Address)
	link.Shortcut = config.GetString("BASE_REDIRECT_SERVER") + link.Shortcut
	c.JSON(http.StatusCreated, gin.H{"data": link})
}
