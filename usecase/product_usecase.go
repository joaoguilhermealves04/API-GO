package usecase

import (
	"Projetos/model"
	"Projetos/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repositoryProduto repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repositoryProduto,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {

	return pu.repository.GetProductos()
}

func (pu *ProductUsecase) AdiconarProduto(produto model.Product) (model.Product, error) {
	produtoId, err := pu.repository.AdicionarProduto(produto)

	if err != nil {
		return model.Product{}, err
	}
	produto.Id = produtoId

	return produto, nil
}

func (pu *ProductUsecase) GebyId(id int) (*model.Product, error) {

	produto, err := pu.GebyId(id)
	if err != nil {
		return nil, err
	}
	return produto, nil
}
