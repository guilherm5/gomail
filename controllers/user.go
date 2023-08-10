package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/database"
	"github.com/guilherm5/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = database.Init()

func NewUser(c *gin.Context) {
	var data models.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("Erro ao ler body da requisição", err)
		c.Status(400)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.Senha), 8)
	if err != nil {
		log.Println("Erro ao gerar hash da senha", err)
		c.Status(500)
		return
	}

	_, err = DB.Exec(`INSERT INTO usuario (nome, email, senha) VALUES ($1, $2, $3)`, &data.Nome, &data.Email, password)
	if err != nil {
		log.Println("Erro ao inserir usuario", err)
		c.Status(400)
		return
	}

	c.Status(201)
}
