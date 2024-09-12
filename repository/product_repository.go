package repository

import (
	"Projetos/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProductos() ([]model.Product, error) {

	query := "SELECT id,product_name,price FROM productos"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	var produtosLita []model.Product
	var produtoObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&produtoObj.Id,
			&produtoObj.Nome,
			&produtoObj.Preco)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		produtosLita = append(produtosLita, produtoObj)
	}
	rows.Close()
	return produtosLita, nil
}

func (pr *ProductRepository) AdicionarProduto(Produto model.Product) (int, error) {

	var id int
	quere, err := pr.connection.Prepare("INSERT INTO producto" +
		"(product_name, price) " + "VALUES($1,$2)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = quere.QueryRow(Produto.Nome, Produto.Preco).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (pr *ProductRepository) GetProductoById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT id,product_name,price FROM productos WHERE id=$1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var produto model.Product

	err = query.QueryRow(id).Scan(
		&produto.Id,
		&produto.Nome,
		&produto.Preco)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &produto, nil
}
