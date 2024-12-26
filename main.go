package main

import (
	"html/template"
	"log"
	"net/http"
	"loja/models"
)

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
	todosOsProdutos := models.BuscaTodosOsProdutos()
	
		
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
