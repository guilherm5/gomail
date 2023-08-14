package controllers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/models"
	"gopkg.in/gomail.v2"
)

func SendMail(c *gin.Context) {
	var data models.Mail
	user := c.GetFloat64("id")
	log.Println(user)
	var pass = os.Getenv("PASS")
	var smtpHost = os.Getenv("smtpHost")
	var From = os.Getenv("Remetente")
	var mail = gomail.NewMessage()

	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("Erro ao ler conteudo do email", err)
		c.Status(400)
		return
	}

	mail.SetHeader("From", From)
	mail.SetHeader("To", data.Destinatario)
	mail.SetHeader("Subject", data.Assunto)
	mail.SetBody("text/plain", data.Conteudo)

	Autentication := gomail.NewDialer(smtpHost, 587, From, pass)
	if err := Autentication.DialAndSend(mail); err != nil {
		log.Println("Erro ao enviar email", err)
		c.Status(400)
		return
	}

	_, err = DB.Exec(`INSERT INTO mail (conteudo, assunto, destinatario, remetente, id_usuario) VALUES ($1, $2, $3, $4, $5)`, &data.Conteudo, &data.Assunto, &data.Destinatario, From, user)
	if err != nil {
		log.Println("Erro ao realizar insert mail", err)
		c.Status(400)
		return
	}

	c.Status(200)
}

func GetMails(c *gin.Context) {
	var jsonData string
	var mailData interface{}

	query, err := DB.Query(`SELECT * FROM  mails`)
	if err != nil {
		log.Println("Erro ao selecionar view", err)
		c.Status(400)
		return
	}

	for query.Next() {
		if err := query.Scan(&jsonData); err != nil {
			log.Println("Erro ao scanear resultado", err)
			c.Status(400)
			return
		}

		if err := json.Unmarshal([]byte(jsonData), &mailData); err != nil {
			log.Println("Erro ao decodificar JSON", err)
			c.Status(500)
			return
		}

	}
	c.JSON(200, mailData)

}
