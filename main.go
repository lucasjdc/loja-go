package main

import (
	"html/template"
	"log"
	"net/http"	
)

type Produto struct {
	Nome		string
	Descricao	string
	Preco		float64
	Quantidade 	int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	
	// Adiciona um log indicando que o servidor está em execução
	log.Println("Servidor iniciado...")
	http.HandleFunc("/", index)
	
	// Trata possíveis erros ao iniciar o servidor
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor:  %v", err)
	}	
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao:"Azul, bem bonita", Preco:39, Quantidade:3},
		{"Tenis","Confortável",89, 3},
		{"Fone","Muito bom", 59, 2},
		{"Lanterna","Ilumina bastante", 200, 2},
	}
	
	temp.ExecuteTemplate(w, "Index", produtos)
}
