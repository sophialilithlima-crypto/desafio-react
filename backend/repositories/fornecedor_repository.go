package repositories

import (
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type FornecedorRepository struct{}

func NewFornecedorRepository() *FornecedorRepository {
	return &FornecedorRepository{}
}


func (r *FornecedorRepository) GetAll() ([]models.Fornecedor, error) {

	rows, err := config.DB.Query(
		`
		SELECT 
			id,
			nome,
			email,
			telefone
		FROM fornecedor
		`,
	)

	if err != nil {
		return nil, err
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
			return nil, err
		}


		fornecedores = append(fornecedores, fornecedor)
	}


	return fornecedores, nil
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
