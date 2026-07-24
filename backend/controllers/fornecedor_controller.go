package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/services"
)

var fornecedorService = services.NewFornecedorService()


func GetFornecedores(c *gin.Context) {

	fornecedores, err := fornecedorService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, fornecedores)
}



func CreateFornecedor(c *gin.Context) {

	var fornecedor models.Fornecedor


	if err := c.ShouldBindJSON(&fornecedor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}


	err := fornecedorService.Create(&fornecedor)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}


	c.JSON(http.StatusCreated, fornecedor)
}



func UpdateFornecedor(c *gin.Context) {

	id := c.Param("id")


	var fornecedor models.Fornecedor


	if err := c.ShouldBindJSON(&fornecedor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}


	err := fornecedorService.Update(id, fornecedor)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Fornecedor atualizado com sucesso",
	})
}



func DeleteFornecedor(c *gin.Context) {

	id := c.Param("id")


	err := fornecedorService.Delete(id)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Fornecedor removido com sucesso",
	})
}