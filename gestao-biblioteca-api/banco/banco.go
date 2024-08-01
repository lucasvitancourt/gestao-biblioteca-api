package banco

import (
	"gerenciamento-biblioteca/modelos"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var BD *gorm.DB

func InicializarBanco() {
	var err error
	BD, err = gorm.Open(sqlite.Open("biblioteca.db"), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar ao banco de dados")
	}

	BD.AutoMigrate(&modelos.Livro{}, &modelos.Autor{}, &modelos.Emprestimo{})
}
