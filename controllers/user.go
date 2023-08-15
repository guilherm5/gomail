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

	if data.Nome == "" || data.Email == "" || data.Senha == "" {
		c.JSON(400, gin.H{
			"Erro ao preencher os dados de cadastro": "Por favor, preencha os campos nome, email e senha",
		})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.Senha), 14)
	if err != nil {
		log.Println("Erro ao gerar hash da senha", err)
		c.Status(500)
		return
	}

	_, err = DB.Exec(`INSERT INTO usuario (nome, email, senha, tipo_usuario) VALUES ($1, $2, $3, $4)`, &data.Nome, &data.Email, password, &data.TipoUsuario)
	if err != nil {
		log.Println("Erro ao inserir usuario", err)
		c.Status(400)
		return
	}

	c.Status(201)
}

func GetUsers(c *gin.Context) {
	var data models.User
	var response []models.User

	query, err := DB.Query(`SELECT id_usuario, nome, email, tipo_usuario FROM usuario`)
	if err != nil {
		log.Println("Erro ao buscar usuarios", err)
		c.Status(400)
		return
	}

	for query.Next() {
		if err := query.Scan(&data.IDUsuario, &data.Nome, &data.Email, &data.TipoUsuario); err != nil {
			log.Println("Erro ao scanear usuarios", err)
			c.Status(400)
			return
		}
		response = append(response, data)
	}
	c.JSON(200, response)
}

func DeleteUsers(c *gin.Context) {
	var data models.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("Erro ao ler body da requisição", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`DELETE FROM usuario WHERE id_usuario = $1`, data.IDUsuario)
}

func UpdateUser(c *gin.Context) {
	var data models.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("Erro ao ler body da requisição", err)
		c.Status(400)
		return
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(data.Senha), 14)

	res, err := DB.Exec(`UPDATE usuario SET nome = $1, email = $2, senha = $3, tipo_usuario = $4 WHERE id_usuario = $5`, &data.Nome, &data.Email, newPassword, &data.TipoUsuario, &data.IDUsuario)
	if err != nil {
		log.Println("Erro ao atualizar usuario", err)
		c.Status(400)
		return
	}
	log.Print(res.RowsAffected())
	c.Status(200)
}

func Test(c *gin.Context) {
	c.JSON(200, "Bem vindo")
}

func Test2(c *gin.Context) {
	c.JSON(200, "Bem vindo")
}
