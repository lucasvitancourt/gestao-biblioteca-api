package roteadores

import (
	"gerenciamento-biblioteca/controladores"

	"github.com/gorilla/mux"
)

func InicializarRoteador() *mux.Router {
	roteador := mux.NewRouter()
	roteador.HandleFunc("/livros", controladores.ObterLivros).Methods("GET")
	roteador.HandleFunc("/livros/{id}", controladores.ObterLivro).Methods("GET")
	roteador.HandleFunc("/livros", controladores.CriarLivro).Methods("POST")
	roteador.HandleFunc("/livros/{id}", controladores.AtualizarLivro).Methods("PUT")
	roteador.HandleFunc("/livros/{id}", controladores.DeletarLivro).Methods("DELETE")
	roteador.HandleFunc("/autores", controladores.ObterAutores).Methods("GET")
	roteador.HandleFunc("/autores/{id}", controladores.ObterAutor).Methods("GET")
	roteador.HandleFunc("/autores", controladores.CriarAutor).Methods("POST")
	roteador.HandleFunc("/autores/{id}", controladores.AtualizarAutor).Methods("PUT")
	roteador.HandleFunc("/autores/{id}", controladores.DeletarAutor).Methods("DELETE")
	roteador.HandleFunc("/emprestimos", controladores.ObterEmprestimos).Methods("GET")
	roteador.HandleFunc("/emprestimos/{id}", controladores.ObterEmprestimo).Methods("GET")
	roteador.HandleFunc("/emprestimos", controladores.CriarEmprestimo).Methods("POST")
	roteador.HandleFunc("/emprestimos/{id}", controladores.AtualizarEmprestimo).Methods("PUT")
	roteador.HandleFunc("/emprestimos/{id}", controladores.DeletarEmprestimo).Methods("DELETE")
	return roteador
}
