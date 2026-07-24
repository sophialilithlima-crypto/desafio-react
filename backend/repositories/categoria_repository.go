package repositories

import (
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type CategoriaRepository struct{}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) GetAll() ([]models.Categoria, error) {

	rows, err := config.DB.Query(
		"SELECT id, nome FROM categoria",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categorias []models.Categoria

	for rows.Next() {

		var categoria models.Categoria

		err := rows.Scan(
			&categoria.ID,
			&categoria.Nome,
		)

		if err != nil {
			return nil, err
		}

		categorias = append(categorias, categoria)
	}

	return categorias, nil
}

func (r *CategoriaRepository) Create(categoria *models.Categoria) error {

	return config.DB.QueryRow(
		"INSERT INTO categoria (nome) VALUES ($1) RETURNING id",
		categoria.Nome,
	).Scan(&categoria.ID)
}

func (r *CategoriaRepository) Update(id string, categoria models.Categoria) error {

	_, err := config.DB.Exec(
		"UPDATE categoria SET nome = $1 WHERE id = $2",
		categoria.Nome,
		id,
	)

	return err
}

func (r *CategoriaRepository) Delete(id string) error {

	_, err := config.DB.Exec(
		"DELETE FROM categoria WHERE id = $1",
		id,
	)

	return err
}