package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/routes"
)

func main() {
	router := gin.Default()

	routes.User(router)
	routes.Login(router)
	routes.Mail(router)

	router.Run(":5555")
}
