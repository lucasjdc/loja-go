package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
    "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Função que conecta ao banco de dados
func conectaComBancoDeDados() *sql.DB {
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

type Produto struct {
	Id			int
	Nome		string
	Descricao	string
	Preco		float64
	Quantidade 	int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {	
	log.Println("Servidor iniciado...")
	http.HandleFunc("/", index)
	
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor:  %v", err)
	}	
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()
	selectDeTodosOsProdutos, err  := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	
	p := Produto{}
	produtos :=  []Produto{}
	
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco  float64
		
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		
		produtos = append(produtos, p)
	}
		
	temp.ExecuteTemplate(w, "Index", produtos)
	db.Close()
}
