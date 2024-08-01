package controladores

import (
	"encoding/json"
	"gerenciamento-biblioteca/banco"
	"gerenciamento-biblioteca/modelos"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ObterAutores(w http.ResponseWriter, r *http.Request) {
	var autores []modelos.Autor
	banco.BD.Preload("Livros").Find(&autores)
	json.NewEncoder(w).Encode(autores)
}

func ObterAutor(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var autor modelos.Autor
	banco.BD.Preload("Livros").First(&autor, id)
	json.NewEncoder(w).Encode(autor)
}

func CriarAutor(w http.ResponseWriter, r *http.Request) {
	var autor modelos.Autor
	json.NewDecoder(r.Body).Decode(&autor)
	banco.BD.Create(&autor)
	json.NewEncoder(w).Encode(autor)
}

func AtualizarAutor(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var autor modelos.Autor
	banco.BD.First(&autor, id)
	json.NewDecoder(r.Body).Decode(&autor)
	banco.BD.Save(&autor)
	json.NewEncoder(w).Encode(autor)
}

func DeletarAutor(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var autor modelos.Autor
	banco.BD.Delete(&autor, id)
	json.NewEncoder(w).Encode(autor)
}
