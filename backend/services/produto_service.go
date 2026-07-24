package services

import (
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/repositories"
)

type ProdutoRepository interface {
	GetAll() ([]models.Produto, error)
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


func (s *ProdutoService) GetAll() ([]models.Produto, error) {

	return s.repository.GetAll()
}


func (s *ProdutoService) Create(produto *models.Produto) error {

	return s.repository.Create(produto)
}


func (s *ProdutoService) Update(id string, produto models.Produto) error {

	return s.repository.Update(id, produto)
}


func (s *ProdutoService) Delete(id string) error {

	return s.repository.Delete(id)
}