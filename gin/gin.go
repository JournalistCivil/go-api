package gin

import (
	"errors"

	"github.com/anon/go-api/internal"
	goapi "github.com/anon/go-api/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	LOCAL      = "local"
	PRODUCTION = "production"
)

type Handler struct {
	UserService goapi.UserService
}

func (h *Handler) ServeHttp(e *internal.Environment) {
	setGinMode(e)

	r := gin.Default()

	r.Static("/storage/public", "storage/public")

	r.Use(setHeaders())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		user, err := h.UserService.FindById(c.Param("id"))
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"message": gorm.ErrRecordNotFound.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data":    user,
			"message": "user retrieved",
		})
	})

	r.Run(":" + e.AppPort)
}

func setGinMode(e *internal.Environment) {
	if e.AppEnv == PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}
}

func setHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}

}
