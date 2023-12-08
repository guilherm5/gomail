package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/models"
	"gopkg.in/gomail.v2"
)

func SendMail(c *gin.Context) {
	var data models.Mail
	user := c.GetFloat64("id")
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

func FileMail(c *gin.Context) {
	//user := c.GetFloat64("id")
	var pass = os.Getenv("PASS")
	var smtpHost = os.Getenv("smtpHost")
	var From = os.Getenv("Remetente")
	var mail = gomail.NewMessage()

	Destinatario := c.PostForm("destinatario")
	Assunto := c.PostForm("assunto")
	Conteudo := c.PostForm("conteudo")

	// tratando envio de arquivos

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Erro ao receber arquivo no body da requisição", err)
		c.Status(400)
		return
	}

	src, err := file.Open()
	if err != nil {
		log.Println("Erro ao abrir arquivo", err)
		c.Status(400)
		return
	}

	src.Seek(0, io.SeekStart)

	mail.SetHeader("From", From)
	mail.SetHeader("To", Destinatario)
	mail.SetHeader("Subject", Assunto)
	mail.SetBody("text/plain", Conteudo)
	mail.Attach(file.Filename, gomail.SetCopyFunc(func(writer io.Writer) error {
		_, err := io.Copy(writer, src)
		return err
	}))

	Autentication := gomail.NewDialer(smtpHost, 587, From, pass)
	if err := Autentication.DialAndSend(mail); err != nil {
		log.Println("Erro ao enviar email", err)
		c.Status(400)
		return
	}

	/*_, err = DB.Exec(`INSERT INTO mail (conteudo, assunto, destinatario, remetente, id_usuario) VALUES ($1, $2, $3, $4, $5)`, &data.Conteudo, &data.Assunto, &data.Destinatario, From, user)
	if err != nil {
		log.Println("Erro ao realizar insert mail", err)
		c.Status(400)
		return
	}*/

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

func GetMailUser(c *gin.Context) {
	var jsonResult json.RawMessage
	user := c.GetFloat64("id")

	query := fmt.Sprintf(`SELECT getMailUser('%v')`, user)
	row := DB.QueryRow(query)

	if err := row.Scan(&jsonResult); err != nil {
		log.Println("Erro ao realizar select na função (bd) get mail usuário", err)
		c.Status(400)
		return
	}
	c.JSON(200, jsonResult)
}
