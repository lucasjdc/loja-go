package db

import (
		"database/sql"
		"github.com/joho/godotenv"
		_ "github.com/lib/pq"
		"os"
		"log"
)

// Função que conecta ao banco de dados
func ConectaComBancoDeDados() *sql.DB {
	// Carrega as variáveis do arquivo .env
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	senha := os.Getenv("DB_PASSWORD")
	conexao := "user=postgres dbname=alura_loja password="+ senha + " host=localhost sslmode=disable"
	db, err :=  sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return db
}
