package services

import (
	"basecode/model"
	"basecode/repository"
	"basecode/web"
	"net/http"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

func (service *ProductServiceImpl) CreateProduct(request web.ProductRequest) (status int, err error) {
	product := model.Product{
		ProductName:  request.ProductName,
		ProductPrice: request.ProductPrice,
		ProductQty:   request.ProductQty,
	}

	err = service.ProductRepository.Save(product)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
