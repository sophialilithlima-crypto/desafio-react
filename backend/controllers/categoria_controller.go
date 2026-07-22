package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

func GetCategorias(c *gin.Context) {

	rows, err := config.DB.Query(
		"SELECT id, nome FROM categoria",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	defer rows.Close()

	var categorias []models.Categoria

	for rows.Next() {

		var categoria models.Categoria

		err := rows.Scan(
			&categoria.ID,
			&categoria.Nome,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})
			return
		}

		categorias = append(categorias, categoria)
	}

	c.JSON(http.StatusOK, categorias)
}

func CreateCategoria(c *gin.Context) {

	var categoria models.Categoria

	// pega os dados enviados pelo JSON
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	// salva no banco
	err := config.DB.QueryRow(
		"INSERT INTO categoria (nome) VALUES ($1) RETURNING id",
		categoria.Nome,
	).Scan(&categoria.ID)

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

	_, err := config.DB.Exec(
		"UPDATE categoria SET nome = $1 WHERE id = $2",
		categoria.Nome,
		id,
	)

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

	_, err := config.DB.Exec(
		"DELETE FROM categoria WHERE id = $1",
		id,
	)

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