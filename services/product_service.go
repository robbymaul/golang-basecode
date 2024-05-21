package services

import "basecode/web"

type ProductService interface {
	CreateProduct(request web.ProductRequest) (status int, err error)
}
