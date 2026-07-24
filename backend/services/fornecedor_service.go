package services

import (
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/repositories"
)


type FornecedorRepository interface {

	GetAll() ([]models.Fornecedor, error)

	Create(fornecedor *models.Fornecedor) error

	Update(id string, fornecedor models.Fornecedor) error

	Delete(id string) error
}



type FornecedorService struct {

	repository FornecedorRepository
}



func NewFornecedorService() *FornecedorService {

	return &FornecedorService{
		repository: repositories.NewFornecedorRepository(),
	}
}



func (s *FornecedorService) GetAll() ([]models.Fornecedor, error) {

	return s.repository.GetAll()
}



func (s *FornecedorService) Create(fornecedor *models.Fornecedor) error {

	return s.repository.Create(fornecedor)
}



func (s *FornecedorService) Update(id string, fornecedor models.Fornecedor) error {

	return s.repository.Update(id, fornecedor)
}



func (s *FornecedorService) Delete(id string) error {

	return s.repository.Delete(id)
}