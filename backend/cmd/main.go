package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/controllers"
)

func main() {

	config.ConnectDatabase()

	router := gin.Default()

	router.GET("/categoria", controllers.GetCategorias)
	router.POST("/categoria", controllers.CreateCategoria)
	router.PUT("/categoria/:id", controllers.UpdateCategoria)
	router.DELETE("/categoria/:id", controllers.DeleteCategoria)

	router.GET("/fornecedor", controllers.GetFornecedores)
	router.POST("/fornecedor", controllers.CreateFornecedor)
	router.PUT("/fornecedor/:id", controllers.UpdateFornecedor)
	router.DELETE("/fornecedor/:id", controllers.DeleteFornecedor)

	router.GET("/produto", controllers.GetProdutos)
	router.POST("/produto", controllers.CreateProduto)
	router.PUT("/produto/:id", controllers.UpdateProduto)
	router.DELETE("/produto/:id", controllers.DeleteProduto)

	fmt.Println("Servidor rodando na porta 8080")

	router.Run(":8080")
}