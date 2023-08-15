package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func User(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.Autentication())
	c.POST("/api/user", controllers.NewUser)
	api.GET("/users", controllers.GetUsers)
	api.DELETE("/delete-user", controllers.DeleteUsers)
	api.PUT("/atualizar-user", controllers.UpdateUser)

	//apenas rota de teste
	api.GET("/test", controllers.Test)
	c.GET("/test2", controllers.Test2)
}
