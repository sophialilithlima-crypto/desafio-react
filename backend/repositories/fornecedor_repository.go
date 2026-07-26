package repositories

import (
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type FornecedorRepository struct{}

func NewFornecedorRepository() *FornecedorRepository {
	return &FornecedorRepository{}
}

func (r *FornecedorRepository) GetAll(page int, limit int, busca string) ([]models.Fornecedor, int, error) {

	offset := (page - 1) * limit

	var total int

	err := config.DB.QueryRow(
		`
		SELECT COUNT(*)
		FROM fornecedor
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
			email,
			telefone
		FROM fornecedor
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

	var fornecedores []models.Fornecedor

	for rows.Next() {

		var fornecedor models.Fornecedor

		err := rows.Scan(
			&fornecedor.ID,
			&fornecedor.Nome,
			&fornecedor.Email,
			&fornecedor.Telefone,
		)

		if err != nil {
			return nil, 0, err
		}

		fornecedores = append(fornecedores, fornecedor)
	}

	return fornecedores, total, nil
}

func (r *FornecedorRepository) GetByID(id string) (models.Fornecedor, error) {

	var fornecedor models.Fornecedor

	err := config.DB.QueryRow(
		`
		SELECT
			id,
			nome,
			email,
			telefone
		FROM fornecedor
		WHERE id = $1
		`,
		id,
	).Scan(
		&fornecedor.ID,
		&fornecedor.Nome,
		&fornecedor.Email,
		&fornecedor.Telefone,
	)

	return fornecedor, err
}

func (r *FornecedorRepository) Create(fornecedor *models.Fornecedor) error {

	return config.DB.QueryRow(
		`
		INSERT INTO fornecedor
		(
			nome,
			email,
			telefone
		)
		VALUES ($1,$2,$3)
		RETURNING id
		`,
		fornecedor.Nome,
		fornecedor.Email,
		fornecedor.Telefone,
	).Scan(&fornecedor.ID)
}

func (r *FornecedorRepository) Update(id string, fornecedor models.Fornecedor) error {

	_, err := config.DB.Exec(
		`
		UPDATE fornecedor
		SET
			nome=$1,
			email=$2,
			telefone=$3
		WHERE id=$4
		`,
		fornecedor.Nome,
		fornecedor.Email,
		fornecedor.Telefone,
		id,
	)

	return err
}

func (r *FornecedorRepository) Delete(id string) error {

	_, err := config.DB.Exec(
		"DELETE FROM fornecedor WHERE id=$1",
		id,
	)

	return err
}