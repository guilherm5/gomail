package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Erro ao carregar variaveis de embiente", err)

	}
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=verify-full", user, database, password, host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados neon", err)

	} else {
		log.Println("Sucesso ao logar no banco de dados")
	}
	return db

}
