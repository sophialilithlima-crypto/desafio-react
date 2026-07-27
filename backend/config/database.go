package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "S0ph14l1l1th")
	dbname := getEnv("DB_NAME", "desafio-react2")

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	database, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()

	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	fmt.Println("Banco conectado com sucesso!")

	DB = database
}

func getEnv(chave string, valorPadrao string) string {

	valor := os.Getenv(chave)

	if valor == "" {
		return valorPadrao
	}

	return valor
}