package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func Mail(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.Autentication())

	api.POST("/mail", controllers.SendMail)
	api.POST("/file-mail", controllers.SendMail)
	api.GET("/mails", controllers.GetMails)
	api.GET("/mail-user", controllers.GetMailUser)
	api.GET("/mail-received", controllers.MailReceived)
}
