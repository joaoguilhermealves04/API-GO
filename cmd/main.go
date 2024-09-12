package main

import (
	"Projetos/Controller" // Importa o pacote Controller, que provavelmente contém as funções responsáveis pelo gerenciamento das requisições HTTP.
	"Projetos/db"
	"Projetos/repository"
	"Projetos/usecase"         // Importa o pacote usecase, que contém a lógica de negócio da aplicação.
	"github.com/gin-gonic/gin" // Importa o pacote gin-gonic/gin, que é um framework web para Go, utilizado para criar o servidor HTTP.
	"log"
)

func main() {
	// Cria uma nova instância do servidor web com as configurações padrão.
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	ProdutoRepository := repository.NewProductRepository(dbConnection)
	// Cria uma nova instância do usecase de produtos. Isso é responsável pela lógica de negócios relacionada a produtos.
	ProdutoUseCase := usecase.NewProductUsecase(ProdutoRepository)

	// Cria um novo controlador para produtos, passando o usecase criado. O controlador será responsável por lidar com as requisições HTTP relacionadas a produtos.
	ProdutoController := Controller.AdiconarProduto(ProdutoUseCase)

	// Define a rota HTTP GET "/produtos" e associa a função GetProdutcts do controlador ProdutoController a esta rota.
	// Quando uma requisição GET é feita para "/produtos", a função GetProdutcts é chamada para tratar a requisição.
	server.GET("/produtos", ProdutoController.GetProdutcts)
	server.POST("/adicionarProduto", ProdutoController.AdicionarProduto)
	server.GET("/produto/:produtoId", ProdutoController.GetProdutcById)

	// Inicia o servidor na porta 8000. O servidor ficará escutando as requisições HTTP nesta porta.
	server.Run(":8000")
}
