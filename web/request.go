package web

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProductRequest struct {
	ProductName  string `json:"product_name" binding:"required"`
	ProductPrice int    `json:"product_price" binding:"required"`
	ProductQty   int    `json:"product_qty" binding:"required"`
}
