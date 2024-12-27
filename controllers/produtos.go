package controllers

import (
	"html/template"
	"net/http"
	"loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	err := temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	if err != nil {
		http.Error(w, "Erro ao renderiza template", http.StatusInternalServerError)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
