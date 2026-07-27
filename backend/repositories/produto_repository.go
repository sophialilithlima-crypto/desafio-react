package repositories

import (
	"fmt"
	"strings"

	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type ProdutoRepository struct{}

func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{}
}

func (r *ProdutoRepository) GetAll(page int, limit int, busca string) ([]models.Produto, int, error) {

	offset := (page - 1) * limit

	var total int

	err := config.DB.QueryRow(
		`
		SELECT COUNT(*)
		FROM produto
		WHERE nome ILIKE $1
		`,
		"%"+busca+"%",
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	rows, err := config.DB.Query(
		`
		SELECT
			id,
			nome,
			sku,
			preco,
			estoque,
			categoria_id,
			criado_em,
			atualizado_em
		FROM produto
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

	var produtos []models.Produto

	for rows.Next() {

		var produto models.Produto

		err := rows.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.SKU,
			&produto.Preco,
			&produto.Estoque,
			&produto.CategoriaID,
			&produto.CriadoEm,
			&produto.AtualizadoEm,
		)

		if err != nil {
			return nil, 0, err
		}

		produtos = append(produtos, produto)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return produtos, total, nil
}

func (r *ProdutoRepository) GetByID(id string) (models.Produto, error) {

	var produto models.Produto

	err := config.DB.QueryRow(
		`
		SELECT
			id,
			nome,
			sku,
			preco,
			estoque,
			categoria_id,
			criado_em,
			atualizado_em
		FROM produto
		WHERE id = $1
		`,
		id,
	).Scan(
		&produto.ID,
		&produto.Nome,
		&produto.SKU,
		&produto.Preco,
		&produto.Estoque,
		&produto.CategoriaID,
		&produto.CriadoEm,
		&produto.AtualizadoEm,
	)

	return produto, err
}

func (r *ProdutoRepository) Create(produto *models.Produto) error {

	err := config.DB.QueryRow(
		`
		INSERT INTO produto
		(
			nome,
			sku,
			preco,
			estoque,
			categoria_id
		)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING id
		`,
		produto.Nome,
		produto.SKU,
		produto.Preco,
		produto.Estoque,
		produto.CategoriaID,
	).Scan(&produto.ID)

	if err != nil {

		if strings.Contains(err.Error(), "duplicate key") ||
			strings.Contains(err.Error(), "duplicar valor da chave") {

			return fmt.Errorf("produto com esse SKU já existe")
		}

		return err
	}

	return nil
}

func (r *ProdutoRepository) Update(id string, produto models.Produto) error {

	_, err := config.DB.Exec(
		`
		UPDATE produto
		SET
			nome=$1,
			sku=$2,
			preco=$3,
			estoque=$4,
			categoria_id=$5,
			atualizado_em=CURRENT_TIMESTAMP
		WHERE id=$6
		`,
		produto.Nome,
		produto.SKU,
		produto.Preco,
		produto.Estoque,
		produto.CategoriaID,
		id,
	)

	if err != nil {

		if strings.Contains(err.Error(), "duplicate key") ||
			strings.Contains(err.Error(), "duplicar valor da chave") {

			return fmt.Errorf("produto com esse SKU já existe")
		}

		return err
	}

	return nil
}

func (r *ProdutoRepository) Delete(id string) error {

	_, err := config.DB.Exec(
		"DELETE FROM produto WHERE id=$1",
		id,
	)

	return err
}