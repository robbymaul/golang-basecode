package repository

import "basecode/model"

type ProductRepository interface {
	Save(product model.Product) error
	Update(product model.Product) error
	FindOne(id int) (model.Product, error)
	FindAll() (product []model.Product, err error)
	Delete(prodcut model.Product) error
}
