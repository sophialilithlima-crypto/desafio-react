package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
	"github.com/sophialilithlima-crypto/desafio-react-backend/controllers"
)


func main() {


	config.ConnectDatabase()



	router := gin.Default()



	// Permitir comunicação com React
	router.Use(cors.New(cors.Config{

		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
		},

	}))





	// =========================
	// CATEGORIA
	// =========================

	router.GET("/categoria", controllers.GetCategorias)
	router.POST("/categoria", controllers.CreateCategoria)
	router.PUT("/categoria/:id", controllers.UpdateCategoria)
	router.DELETE("/categoria/:id", controllers.DeleteCategoria)

	// Rotas no plural para React
	router.GET("/categorias", controllers.GetCategorias)
	router.POST("/categorias", controllers.CreateCategoria)
	router.PUT("/categorias/:id", controllers.UpdateCategoria)
	router.DELETE("/categorias/:id", controllers.DeleteCategoria)






	// =========================
	// FORNECEDOR
	// =========================

	router.GET("/fornecedor", controllers.GetFornecedores)
	router.POST("/fornecedor", controllers.CreateFornecedor)
	router.PUT("/fornecedor/:id", controllers.UpdateFornecedor)
	router.DELETE("/fornecedor/:id", controllers.DeleteFornecedor)

	// Rotas no plural para React
	router.GET("/fornecedores", controllers.GetFornecedores)
	router.POST("/fornecedores", controllers.CreateFornecedor)
	router.PUT("/fornecedores/:id", controllers.UpdateFornecedor)
	router.DELETE("/fornecedores/:id", controllers.DeleteFornecedor)







	// =========================
	// PRODUTO
	// =========================

	router.GET("/produto", controllers.GetProdutos)
	router.POST("/produto", controllers.CreateProduto)
	router.PUT("/produto/:id", controllers.UpdateProduto)
	router.DELETE("/produto/:id", controllers.DeleteProduto)


	// Rotas no plural para React
	router.GET("/produtos", controllers.GetProdutos)
	router.POST("/produtos", controllers.CreateProduto)
	router.PUT("/produtos/:id", controllers.UpdateProduto)
	router.DELETE("/produtos/:id", controllers.DeleteProduto)





	router.GET("/produto/:id", controllers.GetProdutoByID)
	router.GET("/categoria/:id", controllers.GetCategoriaByID)
	router.GET("/fornecedor/:id", controllers.GetFornecedorByID)



	fmt.Println("Servidor rodando na porta 8080")



	router.Run(":8080")

}