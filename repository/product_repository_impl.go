package repository

import (
	"basecode/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}

func (repository *ProductRepositoryImpl) Save(product model.Product) error {
	err := repository.DB.Create(&product).Error

	return err
}

func (repository *ProductRepositoryImpl) Update(product model.Product) error {
	err := repository.DB.Save(product).Error

	return err
}

func (repository *ProductRepositoryImpl) FindOne(id int) (model.Product, error) {
	var product model.Product
	err := repository.DB.Where("id = ?", id).First(&product).Error

	return product, err
}

func (repository *ProductRepositoryImpl) FindAll() (product []model.Product, err error) {
	err = repository.DB.Find(&product).Error

	return product, err
}

func (repository *ProductRepositoryImpl) Delete(product model.Product) error {
	err := repository.DB.Delete(&product).Error

	return err
}
