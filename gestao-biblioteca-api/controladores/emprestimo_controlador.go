package controladores

import (
	"encoding/json"
	"gerenciamento-biblioteca/banco"
	"gerenciamento-biblioteca/modelos"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ObterEmprestimos(w http.ResponseWriter, r *http.Request) {
	var emprestimos []modelos.Emprestimo
	banco.BD.Find(&emprestimos)
	json.NewEncoder(w).Encode(emprestimos)
}

func ObterEmprestimo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var emprestimo modelos.Emprestimo
	banco.BD.First(&emprestimo, id)
	json.NewEncoder(w).Encode(emprestimo)
}

func CriarEmprestimo(w http.ResponseWriter, r *http.Request) {
	var emprestimo modelos.Emprestimo
	json.NewDecoder(r.Body).Decode(&emprestimo)
	banco.BD.Create(&emprestimo)
	json.NewEncoder(w).Encode(emprestimo)
}

func AtualizarEmprestimo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var emprestimo modelos.Emprestimo
	banco.BD.First(&emprestimo, id)
	json.NewDecoder(r.Body).Decode(&emprestimo)
	banco.BD.Save(&emprestimo)
	json.NewEncoder(w).Encode(emprestimo)
}

func DeletarEmprestimo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, _ := strconv.Atoi(parametros["id"])
	var emprestimo modelos.Emprestimo
	banco.BD.Delete(&emprestimo, id)
	json.NewEncoder(w).Encode(emprestimo)
}
