package service

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var s3Session = session.New()

func S3Aws() *session.Session {
	//Variaveis de login para o aws
	region := os.Getenv("REGION")
	key := os.Getenv("KEY")
	secretpass := os.Getenv("SECRETPASS")
	if region == "" || key == "" || secretpass == "" {
		log.Println("informações de login da AWS não pode ser nulas")
	}

	s3Config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secretpass, ""),
	}

	s3Session = session.New(s3Config)
	return session.New(s3Config.WithRegion(region))
}
