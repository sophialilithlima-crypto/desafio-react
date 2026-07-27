package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
	"github.com/sophialilithlima-crypto/desafio-react-backend/services"
)

var categoriaService = services.NewCategoriaService()

func GetCategorias(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil || limit < 1 {
		limit = 10
	}

	busca := c.DefaultQuery("q", "")

	categorias, total, err := categoriaService.GetAll(page, limit, busca)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categorias,
		"error": nil,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func GetCategoriaByID(c *gin.Context) {

	id := c.Param("id")

	categoria, err := categoriaService.GetByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Categoria não encontrada",
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  categoria,
		"error": nil,
		"meta":  nil,
	})
}

func CreateCategoria(c *gin.Context) {

	var categoria models.Categoria

	if err := c.ShouldBindJSON(&categoria); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Dados inválidos",
			"meta":  nil,
		})

		return
	}

	err := categoriaService.Create(&categoria)

	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":  categoria,
		"error": nil,
		"meta":  nil,
	})
}

func UpdateCategoria(c *gin.Context) {

	id := c.Param("id")

	var categoria models.Categoria

	if err := c.ShouldBindJSON(&categoria); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Dados inválidos",
			"meta":  nil,
		})

		return
	}

	err := categoriaService.Update(id, categoria)

	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"error": err.Error(),
			"meta":  nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  categoria,
		"error": nil,
		"meta":  nil,
	})
}

func DeleteCategoria(c *gin.Context) {

	id := c.Param("id")

	err := categoriaService.Delete(id)

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