package main

import (
	"log"
	"net/http"
	"loja/routes"
)

func main() {
	routes.CarregaRotas()	
	log.Println("Servidor iniciado...")
	log.Fatal(http.ListenAndServe(":8080", nil))		
}
