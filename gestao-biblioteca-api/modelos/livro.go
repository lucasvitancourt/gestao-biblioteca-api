package modelos

type Livro struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Titulo     string `json:"titulo"`
	AutorID    uint   `json:"autor_id"`
	Disponivel bool   `json:"disponivel"`
}
