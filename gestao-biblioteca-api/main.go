package main

import (
	"gerenciamento-biblioteca/banco"
	"gerenciamento-biblioteca/roteadores"
	"log"
	"net/http"
)

func main() {
	banco.InicializarBanco()

	roteador := roteadores.InicializarRoteador()
	log.Fatal(http.ListenAndServe(":8080", roteador))
}
