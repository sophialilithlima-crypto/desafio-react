package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/services"
)

var fornecedorService = services.NewFornecedorService()

func GetFornecedores(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil || limit < 1 {
		limit = 10
	}

	busca := c.DefaultQuery("q", "")

	fornecedores, total, err := fornecedorService.GetAll(page, limit, busca)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  fornecedores,
		"error": nil,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func GetFornecedorByID(c *gin.Context) {

	id := c.Param("id")

	fornecedor, err := fornecedorService.GetByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Fornecedor não encontrado",
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  fornecedor,
		"error": nil,
		"meta":  nil,
	})
}

func CreateFornecedor(c *gin.Context) {

	var fornecedor models.Fornecedor

	if err := c.ShouldBindJSON(&fornecedor); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Dados inválidos",
			"meta":  nil,
		})

		return
	}

	err := fornecedorService.Create(&fornecedor)

	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":  fornecedor,
		"error": nil,
		"meta":  nil,
	})
}

func UpdateFornecedor(c *gin.Context) {

	id := c.Param("id")

	var fornecedor models.Fornecedor

	if err := c.ShouldBindJSON(&fornecedor); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Dados inválidos",
			"meta":  nil,
		})

		return
	}

	err := fornecedorService.Update(id, fornecedor)

	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  fornecedor,
		"error": nil,
		"meta":  nil,
	})
}

func DeleteFornecedor(c *gin.Context) {

	id := c.Param("id")

	err := fornecedorService.Delete(id)

	if err != nil {

		c.JSON(http.StatusConflict, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  nil,
		"error": nil,
		"meta":  nil,
	})
}