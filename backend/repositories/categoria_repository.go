package repositories

import (
	"fmt"
	"strings"

	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type CategoriaRepository struct{}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) GetAll(page int, limit int, busca string) ([]models.Categoria, int, error) {

	offset := (page - 1) * limit

	var total int

	buscaParam := "%" + busca + "%"

	err := config.DB.QueryRow(
		`
		SELECT COUNT(*)
		FROM categoria
		WHERE nome ILIKE $1
		`,
		buscaParam,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	rows, err := config.DB.Query(
		`
		SELECT
			id,
			nome
		FROM categoria
		WHERE nome ILIKE $1
		ORDER BY nome
		LIMIT $2 OFFSET $3
		`,
		"%"+busca+"%",
		limit,
		offset,
	)

	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}

		categorias = append(categorias, categoria)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return categorias, total, nil
}

func (r *CategoriaRepository) GetByID(id string) (models.Categoria, error) {

	var categoria models.Categoria

	err := config.DB.QueryRow(
		`
		SELECT
			id,
			nome
		FROM categoria
		WHERE id = $1
		`,
		id,
	).Scan(
		&categoria.ID,
		&categoria.Nome,
	)

	return categoria, err
}

func (r *CategoriaRepository) Create(categoria *models.Categoria) error {

	err := config.DB.QueryRow(
		"INSERT INTO categoria (nome) VALUES ($1) RETURNING id",
		categoria.Nome,
	).Scan(&categoria.ID)

	if err != nil {

		if strings.Contains(err.Error(), "duplicate key") ||
			strings.Contains(err.Error(), "duplicar valor da chave") {

			return fmt.Errorf("categoria já cadastrada")
		}

		return err
	}

	return nil
}

func (r *CategoriaRepository) Update(id string, categoria models.Categoria) error {

	_, err := config.DB.Exec(
		"UPDATE categoria SET nome = $1 WHERE id = $2",
		categoria.Nome,
		id,
	)

	if err != nil {

		if strings.Contains(err.Error(), "duplicate key") ||
			strings.Contains(err.Error(), "duplicar valor da chave") {

			return fmt.Errorf("categoria já cadastrada")
		}

		return err
	}

	return nil
}

func (r *CategoriaRepository) Delete(id string) error {

	_, err := config.DB.Exec(
		"DELETE FROM categoria WHERE id = $1",
		id,
	)

	return err
}

func (r *CategoriaRepository) ExisteProdutoVinculado(id string) (bool, error) {

	var existe bool

	err := config.DB.QueryRow(
		`
		SELECT EXISTS(
			SELECT 1
			FROM produto
			WHERE categoria_id = $1
		)
		`,
		id,
	).Scan(&existe)

	if err != nil {
		return false, err
	}

	return existe, nil
}
