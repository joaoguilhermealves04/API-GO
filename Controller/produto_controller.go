package Controller

import (
	"Projetos/model"
	"Projetos/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProdutoController struct {
	usecase.ProductUsecase
}

func AdiconarProduto(usecase usecase.ProductUsecase) ProdutoController {
	return ProdutoController{
		ProductUsecase: usecase,
	}
}

func (p *ProdutoController) GetProdutcts(ctx *gin.Context) {

	produtos, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, produtos)
}

func (p *ProdutoController) AdicionarProduto(ctx *gin.Context) {

	var produtos model.Product

	err := ctx.BindJSON(&produtos)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.ProductUsecase.AdiconarProduto(produtos)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, insertedProduct)
}

func (p *ProdutoController) GetProdutcById(ctx *gin.Context) {

	Id := ctx.Param("id")

	if Id == "" {

		ctx.JSON(http.StatusNotFound, model.Response.MensagemErroId)
	}
	produtcId, err := strconv.Atoi(Id)
	if Id == "" {

		ctx.JSON(http.StatusNotFound, model.Response.MensagemErroId)
	}

	produto, err := p.ProductUsecase.GebyId(produtcId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response.MensagemErroNoBanco)
	}

	if produto == nil {

		ctx.JSON(http.StatusNotFound, model.Response.MensagemErroNoBanco)
		return
	}

	ctx.JSON(http.StatusOK, produto)
}
