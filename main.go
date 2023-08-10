package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/routes"
)

func main() {
	router := gin.Default()

	routes.User(router)

	router.Run()
}
