package middleware

import (
	"basecode/helper"
	"basecode/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Middleware struct {
	DB *gorm.DB
}

func (middleware *Middleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")
		tokenString := ""

		if len(accessToken) > len("Bearer ") {
			tokenString = accessToken[len("Bearer "):]
		}

		if len(tokenString) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		checkToken, err := helper.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user model.User
		err = middleware.DB.Where("token=?", tokenString).First(&user).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return

		}

		if user.Token != tokenString {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userLogin", checkToken)
		ctx.Next()
	}
}

func (middleware *Middleware) CheckPermission(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userLogin, _ := ctx.Get("userLogin")

		username := userLogin.(jwt.MapClaims)["username"]

		var user model.User
		err := middleware.DB.Where("username=?", username).First(&user).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}

		for _, role := range roles {
			if user.Role == role {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
