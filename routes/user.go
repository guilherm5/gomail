package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/controllers"
)

func User(c *gin.Engine) {
	c.POST("user", controllers.NewUser)
}
