package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/services"
)

var categoriaService = services.NewCategoriaService()

func GetCategorias(c *gin.Context) {

	categorias, err := categoriaService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categorias)
}

func CreateCategoria(c *gin.Context) {

	var categoria models.Categoria

	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	err := categoriaService.Create(&categoria)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, categoria)
}

func UpdateCategoria(c *gin.Context) {

	id := c.Param("id")

	var categoria models.Categoria

	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	err := categoriaService.Update(id, categoria)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Categoria atualizada com sucesso",
	})
}

func DeleteCategoria(c *gin.Context) {

	id := c.Param("id")

	err := categoriaService.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Categoria removida com sucesso",
	})
}