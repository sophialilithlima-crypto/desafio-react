package services

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/repositories"
)

type FornecedorRepository interface {
	GetAll(page int, limit int, busca string) ([]models.Fornecedor, int, error)

	GetByID(id string) (models.Fornecedor, error)

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

func (s *FornecedorService) GetAll(page int, limit int, busca string) ([]models.Fornecedor, int, error) {
	return s.repository.GetAll(page, limit, busca)
}

func (s *FornecedorService) GetByID(id string) (models.Fornecedor, error) {
	return s.repository.GetByID(id)
}

func validarEmail(email string) bool {

	email = strings.TrimSpace(email)

	_, err := mail.ParseAddress(email)

	return err == nil
}

func (s *FornecedorService) validarFornecedor(fornecedor models.Fornecedor) error {

	if strings.TrimSpace(fornecedor.Nome) == "" {
		return fmt.Errorf("nome do fornecedor é obrigatório")
	}

	if !validarEmail(fornecedor.Email) {
		return fmt.Errorf("email inválido")
	}

	return nil
}

func (s *FornecedorService) Create(fornecedor *models.Fornecedor) error {

	if fornecedor == nil {
		return fmt.Errorf("fornecedor não pode ser nulo")
	}

	err := s.validarFornecedor(*fornecedor)

	if err != nil {
		return err
	}

	return s.repository.Create(fornecedor)
}

func (s *FornecedorService) Update(id string, fornecedor models.Fornecedor) error {

	err := s.validarFornecedor(fornecedor)

	if err != nil {
		return err
	}

	return s.repository.Update(id, fornecedor)
}

func (s *FornecedorService) Delete(id string) error {
	return s.repository.Delete(id)
}