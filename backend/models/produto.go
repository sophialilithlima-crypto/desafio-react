package models

type Produto struct {
	ID            int     `json:"id"`
	Nome          string  `json:"nome"`
	Preco         float64 `json:"preco"`
	CategoriaID   int     `json:"categoria_id"`
	FornecedorID  int     `json:"fornecedor_id"`
}