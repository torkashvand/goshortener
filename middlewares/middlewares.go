package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//DBMiddlware add db connection to the request
func DBMiddlware(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}

}
