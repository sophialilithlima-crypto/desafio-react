package services

import (
	"fmt"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/repositories"
	"github.com/sophialilithlima-crypto/desafio-react-backend/utils"
)



type ProdutoRepository interface {

	GetAll(page int, limit int, busca string) ([]models.Produto, int, error)

	GetByID(id string) (models.Produto, error)

	Create(produto *models.Produto) error

	Update(id string, produto models.Produto) error

	Delete(id string) error
}





type ProdutoService struct {

	repository ProdutoRepository
}





func NewProdutoService() *ProdutoService {

	return &ProdutoService{
		repository: repositories.NewProdutoRepository(),
	}
}





func (s *ProdutoService) GetAll(page int, limit int, busca string) ([]models.Produto, int, error) {

	return s.repository.GetAll(page, limit, busca)

}





func (s *ProdutoService) GetByID(id string) (models.Produto, error) {

	return s.repository.GetByID(id)

}





func (s *ProdutoService) Create(produto *models.Produto) error {


	if !utils.ValidarTexto(produto.Nome) {

		return fmt.Errorf("nome do produto é obrigatório")

	}



	if !utils.ValidarTexto(produto.SKU) {

		return fmt.Errorf("sku do produto é obrigatório")

	}



	if produto.Preco <= 0 {

		return fmt.Errorf("preço deve ser maior que zero")

	}



	if produto.Estoque < 0 {

		return fmt.Errorf("estoque não pode ser negativo")

	}



	return s.repository.Create(produto)

}





func (s *ProdutoService) Update(id string, produto models.Produto) error {


	if !utils.ValidarTexto(produto.Nome) {

		return fmt.Errorf("nome do produto é obrigatório")

	}



	if !utils.ValidarTexto(produto.SKU) {

		return fmt.Errorf("sku do produto é obrigatório")

	}



	if produto.Preco <= 0 {

		return fmt.Errorf("preço deve ser maior que zero")

	}



	if produto.Estoque < 0 {

		return fmt.Errorf("estoque não pode ser negativo")

	}



	return s.repository.Update(id, produto)

}





func (s *ProdutoService) Delete(id string) error {

	return s.repository.Delete(id)

}