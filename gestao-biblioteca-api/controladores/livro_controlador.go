package controladores

import (
	"encoding/json"
	"gerenciamento-biblioteca/banco"
	"gerenciamento-biblioteca/modelos"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ObterLivros(w http.ResponseWriter, r *http.Request) {
	var livros []modelos.Livro
	banco.BD.Find(&livros)
	json.NewEncoder(w).Encode(livros)
}

func ObterLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var livro modelos.Livro
	banco.BD.First(&livro, id)
	json.NewEncoder(w).Encode(livro)
}

func CriarLivro(w http.ResponseWriter, r *http.Request) {
	var livro modelos.Livro
	json.NewDecoder(r.Body).Decode(&livro)
	banco.BD.Create(&livro)
	json.NewEncoder(w).Encode(livro)
}

func AtualizarLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var livro modelos.Livro
	banco.BD.First(&livro, id)
	json.NewDecoder(r.Body).Decode(&livro)
	banco.BD.Save(&livro)
	json.NewEncoder(w).Encode(livro)
}

func DeletarLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var livro modelos.Livro
	banco.BD.Delete(&livro, id)
	json.NewEncoder(w).Encode(livro)
}
