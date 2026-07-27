package services

import (
	"fmt"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/repositories"
	"github.com/sophialilithlima-crypto/desafio-react-backend/utils"
)

type CategoriaRepository interface {
	GetAll(page int, limit int, busca string) ([]models.Categoria, int, error)

	GetByID(id string) (models.Categoria, error)

	Create(categoria *models.Categoria) error

	Update(id string, categoria models.Categoria) error

	Delete(id string) error

	ExisteProdutoVinculado(id string) (bool, error)
}

type CategoriaService struct {
	repository CategoriaRepository
}

func NewCategoriaService() *CategoriaService {
	return &CategoriaService{
		repository: repositories.NewCategoriaRepository(),
	}
}

func (s *CategoriaService) GetAll(page int, limit int, busca string) ([]models.Categoria, int, error) {
	return s.repository.GetAll(page, limit, busca)
}

func (s *CategoriaService) GetByID(id string) (models.Categoria, error) {
	return s.repository.GetByID(id)
}

func (s *CategoriaService) Create(categoria *models.Categoria) error {

	if categoria == nil {
		return fmt.Errorf("categoria não pode ser nula")
	}

	if !utils.ValidarTexto(categoria.Nome) {
		return fmt.Errorf("nome da categoria é obrigatório")
	}

	return s.repository.Create(categoria)
}

func (s *CategoriaService) Update(id string, categoria models.Categoria) error {

	if !utils.ValidarTexto(categoria.Nome) {
		return fmt.Errorf("nome da categoria é obrigatório")
	}

	return s.repository.Update(id, categoria)
}

func (s *CategoriaService) Delete(id string) error {

	existeProduto, err := s.repository.ExisteProdutoVinculado(id)

	if err != nil {
		return err
	}

	if existeProduto {
		return fmt.Errorf(
			"não é possível excluir a categoria pois existem produtos vinculados",
		)
	}

	return s.repository.Delete(id)
}