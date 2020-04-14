package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/torkashvand/goshortener/cache"
	"github.com/torkashvand/goshortener/models"
)

//Redirect redirect to shortend address
func Redirect(c *gin.Context) {

	cache.InitializeRedis()
	shortcut := c.Param("shortcut")
	address, _ := cache.GetValue(shortcut)

	if address != nil {
		c.Redirect(http.StatusMovedPermanently, address.(string))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var link models.Link

	if err := db.Where("shortcut = ?", shortcut).First(&link).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	cache.SetValue(shortcut, link.Address)

	c.Redirect(http.StatusMovedPermanently, link.Address)
	c.Abort()
}
