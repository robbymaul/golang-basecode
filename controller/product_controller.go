package controller

import (
	"basecode/services"
	"basecode/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (controller *ProductController) CreateProduct(c *gin.Context) {
	request := new(web.ProductRequest)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	status, err := controller.ProductService.CreateProduct(*request)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{"message": "success create product"})

}
