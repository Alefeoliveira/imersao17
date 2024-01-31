package service

import "github.com/Alefeoliveira/imersao17/goapi/internal/db"

type ProductService struct {
	ProductDB db.ProductDB
}

func NewProductService(productDb db.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDb}
}
