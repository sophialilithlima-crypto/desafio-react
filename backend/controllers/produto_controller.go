package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/services"
)

var produtoService = services.NewProdutoService()

func GetProdutos(c *gin.Context) {

	produtos, err := produtoService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, produtos)
}


func CreateProduto(c *gin.Context) {

	var produto models.Produto

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}


	err := produtoService.Create(&produto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}


	c.JSON(http.StatusCreated, produto)
}


func UpdateProduto(c *gin.Context) {

	id := c.Param("id")

	var produto models.Produto

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}


	err := produtoService.Update(id, produto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Produto atualizado com sucesso",
	})
}


func DeleteProduto(c *gin.Context) {

	id := c.Param("id")

	err := produtoService.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Produto removido com sucesso",
	})
}