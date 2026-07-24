package services

import (
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/repositories"
)

type CategoriaRepository interface {
	GetAll() ([]models.Categoria, error)
	Create(categoria *models.Categoria) error
	Update(id string, categoria models.Categoria) error
	Delete(id string) error
}

type CategoriaService struct {
	repository CategoriaRepository
}

func NewCategoriaService() *CategoriaService {
	return &CategoriaService{
		repository: repositories.NewCategoriaRepository(),
	}
}

func (s *CategoriaService) GetAll() ([]models.Categoria, error) {
	return s.repository.GetAll()
}

func (s *CategoriaService) Create(categoria *models.Categoria) error {
	return s.repository.Create(categoria)
}

func (s *CategoriaService) Update(id string, categoria models.Categoria) error {
	return s.repository.Update(id, categoria)
}

func (s *CategoriaService) Delete(id string) error {
	return s.repository.Delete(id)
}