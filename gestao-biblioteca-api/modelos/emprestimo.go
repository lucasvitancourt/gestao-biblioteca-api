package modelos

import "time"

type Emprestimo struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	LivroID        uint      `json:"livro_id"`
	Usuario        string    `json:"usuario"`
	DataEmprestimo time.Time `json:"data_emprestimo"`
	DataDevolucao  time.Time `json:"data_devolucao"`
}
