package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type fornecedorRepositoryMock struct{}

func (m *fornecedorRepositoryMock) GetAll(page int, limit int, busca string) ([]models.Fornecedor, int, error) {
	return []models.Fornecedor{}, 0, nil
}

func (m *fornecedorRepositoryMock) GetByID(id string) (models.Fornecedor, error) {
	return models.Fornecedor{}, nil
}

func (m *fornecedorRepositoryMock) Create(fornecedor *models.Fornecedor) error {
	return nil
}

func (m *fornecedorRepositoryMock) Update(id string, fornecedor models.Fornecedor) error {
	return nil
}

func (m *fornecedorRepositoryMock) Delete(id string) error {
	return nil
}


func TestFornecedorCreateSucesso(t *testing.T) {

	service := FornecedorService{
		repository: &fornecedorRepositoryMock{},
	}

	fornecedor := models.Fornecedor{
		Nome:     "Fornecedor Teste",
		Email:    "teste@email.com",
		Telefone: "81999999999",
	}

	err := service.Create(&fornecedor)

	assert.Nil(t, err)
}


func TestFornecedorCreateNomeObrigatorio(t *testing.T) {

	service := FornecedorService{
		repository: &fornecedorRepositoryMock{},
	}

	fornecedor := models.Fornecedor{
		Nome:  "",
		Email: "teste@email.com",
	}

	err := service.Create(&fornecedor)

	assert.NotNil(t, err)
	assert.Equal(t, "nome do fornecedor é obrigatório", err.Error())
}


func TestFornecedorCreateEmailInvalido(t *testing.T) {

	service := FornecedorService{
		repository: &fornecedorRepositoryMock{},
	}

	fornecedor := models.Fornecedor{
		Nome:  "Fornecedor Teste",
		Email: "email-invalido",
	}

	err := service.Create(&fornecedor)

	assert.NotNil(t, err)
	assert.Equal(t, "email inválido", err.Error())
}


type fornecedorRepositoryErro struct{}

func (m *fornecedorRepositoryErro) GetAll(page int, limit int, busca string) ([]models.Fornecedor, int, error) {
	return nil, 0, errors.New("erro")
}

func (m *fornecedorRepositoryErro) GetByID(id string) (models.Fornecedor, error) {
	return models.Fornecedor{}, errors.New("erro")
}

func (m *fornecedorRepositoryErro) Create(fornecedor *models.Fornecedor) error {
	return errors.New("erro")
}

func (m *fornecedorRepositoryErro) Update(id string, fornecedor models.Fornecedor) error {
	return errors.New("erro")
}

func (m *fornecedorRepositoryErro) Delete(id string) error {
	return errors.New("erro")
}


func TestFornecedorRepositoryErro(t *testing.T) {

	service := FornecedorService{
		repository: &fornecedorRepositoryErro{},
	}

	fornecedor := models.Fornecedor{
		Nome:  "Fornecedor Teste",
		Email: "teste@email.com",
	}

	err := service.Create(&fornecedor)

	assert.NotNil(t, err)
	assert.Equal(t, "erro", err.Error())
}