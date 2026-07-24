package repositories

import (
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type ProdutoRepository struct{}

func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{}
}


func (r *ProdutoRepository) GetAll() ([]models.Produto, error) {

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
		`,
	)

	if err != nil {
		return nil, err
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
			return nil, err
		}


		produtos = append(produtos, produto)
	}


	return produtos, nil
}



func (r *ProdutoRepository) Create(produto *models.Produto) error {


	return config.DB.QueryRow(
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


	return err
}



func (r *ProdutoRepository) Delete(id string) error {


	_, err := config.DB.Exec(
		"DELETE FROM produto WHERE id=$1",
		id,
	)


	return err
}