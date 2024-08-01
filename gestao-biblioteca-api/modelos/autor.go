package modelos

type Autor struct {
	ID     uint    `json:"id" gorm:"primary_key"`
	Nome   string  `json:"nome"`
	Livros []Livro `json:"livros" gorm:"foreignKey:AutorID"`
}
