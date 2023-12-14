package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/guilherm5/models"
	"github.com/guilherm5/service"

	"gopkg.in/gomail.v2"
)

func SendMail(c *gin.Context) {
	var data models.Mail
	user := c.GetFloat64("id")
	var pass = os.Getenv("PASS")
	var smtpHost = os.Getenv("smtpHost")
	var From = os.Getenv("Remetente")
	var mail = gomail.NewMessage()

	if c.Request.URL.Path == "/api/mail" {
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
	} else if c.Request.URL.Path == "/api/file-mail" {
		strFile, err := uuid.NewV4()
		if err != nil {
			log.Println("Erro ao gerar uuid foto", err)
			c.Status(400)
			return
		}
		service := service.S3Aws()
		Destinatario := c.PostForm("destinatario")
		Assunto := c.PostForm("assunto")
		Conteudo := c.PostForm("conteudo")
		file, err := c.FormFile("file")
		if err != nil {
			log.Println("Erro ao receber arquivo no body da requisição", err)
			c.Status(400)
			return
		}
		if file == nil {
			log.Println("Arquivo não pode ser nulo")
			c.Status(400)
			return
		}

		src, err := file.Open()
		if err != nil {
			log.Println("Erro ao abrir arquivo", err)
			c.Status(400)
			return
		}
		_, err = io.ReadAll(src)
		if err != nil {
			log.Println("Erro ao ler arquivo", err)
			c.Status(400)
			return
		}
		src.Seek(0, io.SeekStart)
		uploader := s3manager.NewUploader(service)
		contentType := file.Header.Get("Content-Type")
		input := &s3manager.UploadInput{
			Bucket:             aws.String("gomail-go"),
			Key:                aws.String("files-mail/" + strFile.String()),
			Body:               src,
			ContentType:        &contentType,
			ContentDisposition: aws.String("inline"),
		}

		_, err = uploader.UploadWithContext(context.Background(), input)
		if err != nil {
			log.Println("Erro ao realizar uplaod da imagem no bucket s3", err)
			c.Status(101)
		}
		linkPost := fmt.Sprintf("https://frienlinkfotos.s3.amazonaws.com/%s", strFile)

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

		_, err = DB.Exec(`INSERT INTO mail (conteudo, assunto, destinatario, remetente, id_usuario, caminho_arquivo, uuid_mail) VALUES ($1, $2, $3, $4, $5, $6, $7)`, Conteudo, Assunto, Destinatario, From, user, linkPost, strFile)
		if err != nil {
			log.Println("Erro ao realizar insert mail", err)
			c.Status(400)
			return
		}
		c.Status(200)
	}
}

func MailReceived(c *gin.Context) {
	var server = os.Getenv("server")
	var username = os.Getenv("Remetente")
	var password = os.Getenv("PASS")

	// Conectar ao servidor IMAP
	im, err := client.DialTLS(fmt.Sprintf("%s:993", server), nil)
	if err != nil {
		log.Fatal("falha ao conectar no servidor imap", err)
		c.Status(400)
		return
	}
	defer im.Logout()

	// Autenticar com o servidor IMAP
	if err := im.Login(username, password); err != nil {
		log.Fatal("falha ao conectar no servidor imap", err)
		c.Status(400)
		return
	}

	// Selecionar a caixa de correio
	mbox, err := im.Select("INBOX", false)
	if err != nil {
		log.Fatal("erro ao obter emails", err)
		c.Status(400)
		return
	}

	// Listar os IDs dos e-mails na caixa de correio
	from := uint32(1)
	to := mbox.Messages
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	go func() {
		if err := im.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages); err != nil {
			log.Fatal(err)
		}
	}()
	dataAtual := time.Now().Format("02-Jan-2006")
	for msg := range messages {
		dataMensagem := msg.Envelope.Date.Format("02-Jan-2006")
		if dataMensagem == dataAtual {
			c.JSON(200, gin.H{
				"Assunto:": msg.Envelope.Subject,
				"De:":      msg.Envelope.From[0],
				"Data:":    msg.Envelope.Date.Format("02-Jan-2006"),
			})
		}

	}
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
