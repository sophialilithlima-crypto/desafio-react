package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type produtoRepositoryMock struct{}

func (m *produtoRepositoryMock) GetAll(page int, limit int, busca string) ([]models.Produto, int, error) {
	return []models.Produto{}, 0, nil
}

func (m *produtoRepositoryMock) GetByID(id string) (models.Produto, error) {
	return models.Produto{}, nil
}

func (m *produtoRepositoryMock) Create(produto *models.Produto) error {
	return nil
}

func (m *produtoRepositoryMock) Update(id string, produto models.Produto) error {
	return nil
}

func (m *produtoRepositoryMock) Delete(id string) error {
	return nil
}


func TestProdutoCreateSucesso(t *testing.T) {

	service := ProdutoService{
		repository: &produtoRepositoryMock{},
	}

	produto := models.Produto{
		Nome:       "Notebook",
		SKU:        "NOTE123",
		Preco:      2500,
		Estoque:    10,
		CategoriaID: 1,
	}

	err := service.Create(&produto)

	assert.Nil(t, err)
}


func TestProdutoCreateNomeObrigatorio(t *testing.T) {

	service := ProdutoService{
		repository: &produtoRepositoryMock{},
	}

	produto := models.Produto{
		Nome:       "",
		SKU:        "NOTE123",
		Preco:      2500,
		Estoque:    10,
		CategoriaID: 1,
	}

	err := service.Create(&produto)

	assert.NotNil(t, err)
	assert.Equal(t, "nome do produto é obrigatório", err.Error())
}


func TestProdutoCreateSKUObrigatorio(t *testing.T) {

	service := ProdutoService{
		repository: &produtoRepositoryMock{},
	}

	produto := models.Produto{
		Nome:       "Notebook",
		SKU:        "",
		Preco:      2500,
		Estoque:    10,
		CategoriaID: 1,
	}

	err := service.Create(&produto)

	assert.NotNil(t, err)
	assert.Equal(t, "sku do produto é obrigatório", err.Error())
}


func TestProdutoCreatePrecoInvalido(t *testing.T) {

	service := ProdutoService{
		repository: &produtoRepositoryMock{},
	}

	produto := models.Produto{
		Nome:       "Notebook",
		SKU:        "NOTE123",
		Preco:      0,
		Estoque:    10,
		CategoriaID: 1,
	}

	err := service.Create(&produto)

	assert.NotNil(t, err)
	assert.Equal(t, "preço deve ser maior que zero", err.Error())
}


func TestProdutoCreateEstoqueNegativo(t *testing.T) {

	service := ProdutoService{
		repository: &produtoRepositoryMock{},
	}

	produto := models.Produto{
		Nome:       "Notebook",
		SKU:        "NOTE123",
		Preco:      2500,
		Estoque:    -1,
		CategoriaID: 1,
	}

	err := service.Create(&produto)

	assert.NotNil(t, err)
	assert.Equal(t, "estoque não pode ser negativo", err.Error())
}



type produtoRepositoryErro struct{}

func (m *produtoRepositoryErro) GetAll(page int, limit int, busca string) ([]models.Produto, int, error) {
	return nil, 0, errors.New("erro")
}

func (m *produtoRepositoryErro) GetByID(id string) (models.Produto, error) {
	return models.Produto{}, errors.New("erro")
}

func (m *produtoRepositoryErro) Create(produto *models.Produto) error {
	return errors.New("erro")
}

func (m *produtoRepositoryErro) Update(id string, produto models.Produto) error {
	return errors.New("erro")
}

func (m *produtoRepositoryErro) Delete(id string) error {
	return errors.New("erro")
}


func TestProdutoRepositoryErro(t *testing.T) {

	service := ProdutoService{
		repository: &produtoRepositoryErro{},
	}

	produto := models.Produto{
		Nome:       "Notebook",
		SKU:        "NOTE123",
		Preco:      2500,
		Estoque:    10,
		CategoriaID: 1,
	}

	err := service.Create(&produto)

	assert.NotNil(t, err)
	assert.Equal(t, "erro", err.Error())
}