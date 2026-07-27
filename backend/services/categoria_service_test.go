package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

type categoriaRepositoryMock struct{}

func (m *categoriaRepositoryMock) GetAll(page int, limit int, busca string) ([]models.Categoria, int, error) {
	return []models.Categoria{}, 0, nil
}

func (m *categoriaRepositoryMock) GetByID(id string) (models.Categoria, error) {
	return models.Categoria{}, nil
}

func (m *categoriaRepositoryMock) Create(categoria *models.Categoria) error {
	return nil
}

func (m *categoriaRepositoryMock) Update(id string, categoria models.Categoria) error {
	return nil
}

func (m *categoriaRepositoryMock) Delete(id string) error {
	return nil
}

func (m *categoriaRepositoryMock) ExisteProdutoVinculado(id string) (bool, error) {
	return false, nil
}

func TestCategoriaCreateSucesso(t *testing.T) {

	service := CategoriaService{
		repository: &categoriaRepositoryMock{},
	}

	categoria := models.Categoria{
		Nome: "Eletrônicos",
	}

	err := service.Create(&categoria)

	assert.Nil(t, err)
}

func TestCategoriaCreateNomeObrigatorio(t *testing.T) {

	service := CategoriaService{
		repository: &categoriaRepositoryMock{},
	}

	categoria := models.Categoria{
		Nome: "",
	}

	err := service.Create(&categoria)

	assert.NotNil(t, err)
	assert.Equal(t, "nome da categoria é obrigatório", err.Error())
}

func TestCategoriaDeleteSemProdutoVinculado(t *testing.T) {

	service := CategoriaService{
		repository: &categoriaRepositoryMock{},
	}

	err := service.Delete("1")

	assert.Nil(t, err)
}


type categoriaRepositoryErro struct{}

func (m *categoriaRepositoryErro) GetAll(page int, limit int, busca string) ([]models.Categoria, int, error) {
	return nil, 0, errors.New("erro")
}

func (m *categoriaRepositoryErro) GetByID(id string) (models.Categoria, error) {
	return models.Categoria{}, errors.New("erro")
}

func (m *categoriaRepositoryErro) Create(categoria *models.Categoria) error {
	return errors.New("erro")
}

func (m *categoriaRepositoryErro) Update(id string, categoria models.Categoria) error {
	return errors.New("erro")
}

func (m *categoriaRepositoryErro) Delete(id string) error {
	return errors.New("erro")
}

func (m *categoriaRepositoryErro) ExisteProdutoVinculado(id string) (bool, error) {
	return false, errors.New("erro")
}

func TestCategoriaRepositoryErro(t *testing.T) {

	service := CategoriaService{
		repository: &categoriaRepositoryErro{},
	}

	categoria := models.Categoria{
		Nome: "Teste",
	}

	err := service.Create(&categoria)

	assert.NotNil(t, err)
	assert.Equal(t, "erro", err.Error())
}