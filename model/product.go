package model

type Product struct {
	Id    int     `json:"id_Produto"`
	Nome  string  `json:"Nome"`
	Preco float64 `json:"preco"`
}
