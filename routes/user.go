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
	//rotas para adm
	api.GET("/users", controllers.GetUsers)
	api.DELETE("/delete-user", controllers.DeleteUsers)
	api.PUT("/atualizar-user", controllers.UpdateUser)

	//crud para usuarios
	api.GET("/my-user", controllers.GetMyUser)
	api.PUT("/update-secret-my-user", controllers.UpdateMyUser)
	api.PUT("/update-name-my-user", controllers.UpdateMyUser)
	api.DELETE("/delete-my-user", controllers.DeleteMyUser)

}
