package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/services"
)

var produtoService = services.NewProdutoService()

func GetProdutos(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil || limit < 1 {
		limit = 10
	}

	busca := c.DefaultQuery("q", "")

	produtos, total, err := produtoService.GetAll(page, limit, busca)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  produtos,
		"error": nil,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func GetProdutoByID(c *gin.Context) {

	id := c.Param("id")

	produto, err := produtoService.GetByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Produto não encontrado",
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  produto,
		"error": nil,
		"meta":  nil,
	})
}

func CreateProduto(c *gin.Context) {

	var produto models.Produto

	if err := c.ShouldBindJSON(&produto); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Dados inválidos",
			"meta":  nil,
		})

		return
	}

	err := produtoService.Create(&produto)

	if err != nil {

		if strings.Contains(strings.ToLower(err.Error()), "sku") {

			c.JSON(http.StatusConflict, gin.H{
				"data":  nil,
				"error": err.Error(),
				"meta":  nil,
			})

			return
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":  produto,
		"error": nil,
		"meta":  nil,
	})
}

func UpdateProduto(c *gin.Context) {

	id := c.Param("id")

	var produto models.Produto

	if err := c.ShouldBindJSON(&produto); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Dados inválidos",
			"meta":  nil,
		})

		return
	}

	err := produtoService.Update(id, produto)

	if err != nil {

		if strings.Contains(strings.ToLower(err.Error()), "sku") {

			c.JSON(http.StatusConflict, gin.H{
				"data":  nil,
				"error": err.Error(),
				"meta":  nil,
			})

			return
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  produto,
		"error": nil,
		"meta":  nil,
	})
}

func DeleteProduto(c *gin.Context) {

	id := c.Param("id")

	err := produtoService.Delete(id)

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