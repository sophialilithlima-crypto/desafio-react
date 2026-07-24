package models

import "time"

type Produto struct {
	ID          int       `json:"id"`
	Nome        string    `json:"nome"`
	SKU         string    `json:"sku"`
	Preco       float64   `json:"preco"`
	Estoque     int       `json:"estoque"`
	CategoriaID int       `json:"categoria_id"`
	CriadoEm    time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
}