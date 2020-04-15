import (
	"github.com/gin-gonic/gin"
)

func Db(c *gin.Context) {
	c.Set("db", db)
	c.Next()
}