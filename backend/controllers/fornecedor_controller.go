package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)

func GetFornecedores(c *gin.Context) {

	rows, err := config.DB.Query(
		"SELECT id, nome, email FROM fornecedor",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	defer rows.Close()

	var fornecedores []models.Fornecedor

	for rows.Next() {

		var fornecedor models.Fornecedor

		err := rows.Scan(
			&fornecedor.ID,
			&fornecedor.Nome,
			&fornecedor.Email,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})
			return
		}

		fornecedores = append(fornecedores, fornecedor)
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


	err := config.DB.QueryRow(
		"INSERT INTO fornecedor (nome, email) VALUES ($1, $2) RETURNING id",
		fornecedor.Nome,
		fornecedor.Email,
	).Scan(&fornecedor.ID)


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


	_, err := config.DB.Exec(
		"UPDATE fornecedor SET nome=$1, email=$2 WHERE id=$3",
		fornecedor.Nome,
		fornecedor.Email,
		id,
	)


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


	_, err := config.DB.Exec(
		"DELETE FROM fornecedor WHERE id=$1",
		id,
	)


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