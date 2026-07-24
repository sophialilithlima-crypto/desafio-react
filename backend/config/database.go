package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {

	connection := "host=localhost port=5432 user=postgres password=S0ph14l1l1th dbname=desafio-react2 sslmode=disable"

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