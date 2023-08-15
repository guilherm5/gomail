package controllers

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Erro ao carregar variaveis de embiente", err)
		c.Status(400)
		return
	}
	var credentials models.User
	var login = models.User{}
	secret := os.Getenv("TOKEN")

	err = c.ShouldBindJSON(&credentials)
	if err != nil {
		log.Println("Erro ao ler body da requisição", err)
		c.Status(400)
		return
	}

	query, err := DB.Query(`SELECT * FROM usuario WHERE email = $1`, credentials.Email)
	if err != nil {
		log.Println("Erro ao buscar usuario para realizar login", err)
		c.Status(400)
		return
	}

	for query.Next() {
		err = query.Scan(&login.IDUsuario, &login.Nome, &login.Email, &login.Senha, &login.TipoUsuario)
		if err != nil {
			log.Println("Erro ao buscar usuario para realizar login", err)
			c.Status(400)
			return
		}
		c.Next()
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Senha), []byte(credentials.Senha))
	if err != nil {
		log.Println("Senha inválida", err)
		c.JSON(401, gin.H{
			"Erro": "Senha inválida",
		})
		return
	}

	claims := jwt.MapClaims{
		"ID":   login.IDUsuario,
		"Nome": login.Nome,
		"Perm": login.TipoUsuario,
		"exp":  time.Hour,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Eroo ao gerar JWT", err)
		c.Status(500)
		return
	}

	c.JSON(200, tokenString)
}
