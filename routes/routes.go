package routes

import (
	"net/http"
	"loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
