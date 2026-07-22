package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/models"
)


func GetProdutos(c *gin.Context) {

	rows, err := config.DB.Query(
		"SELECT id, nome, preco, categoria_id, fornecedor_id FROM produto",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	defer rows.Close()

	var produtos []models.Produto

	for rows.Next() {

		var produto models.Produto

		err := rows.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Preco,
			&produto.CategoriaID,
			&produto.FornecedorID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})
			return
		}

		produtos = append(produtos, produto)
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


	err := config.DB.QueryRow(
		"INSERT INTO produto (nome, preco, categoria_id, fornecedor_id) VALUES ($1,$2,$3,$4) RETURNING id",
		produto.Nome,
		produto.Preco,
		produto.CategoriaID,
		produto.FornecedorID,
	).Scan(&produto.ID)


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

	_, err := config.DB.Exec(
		`UPDATE produto
		 SET nome=$1, preco=$2, categoria_id=$3, fornecedor_id=$4
		 WHERE id=$5`,
		produto.Nome,
		produto.Preco,
		produto.CategoriaID,
		produto.FornecedorID,
		id,
	)

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

	_, err := config.DB.Exec(
		"DELETE FROM produto WHERE id=$1",
		id,
	)

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