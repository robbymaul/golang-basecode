package services

import (
	"basecode/helper"
	"basecode/model"
	"basecode/repository"
	"basecode/web"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{UserRepository: userRepository}
}

func (service *AuthServiceImpl) Register(request *web.RegisterUserRequest) (status int, err error) {
	passwordHash, err := helper.HashingPassword(request.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user := model.User{
		Username: request.Username,
		Password: passwordHash,
		Role:     request.Role,
	}

	err = service.UserRepository.Save(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (service *AuthServiceImpl) Login(request web.LoginUserRequest) (response web.LoginUserResponse, status int, err error) {

	user, err := service.UserRepository.FindOneByUsername(request.Username)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	validPassword := helper.ComparePassword(request.Password, user.Password)
	if !validPassword {
		return response, http.StatusUnauthorized, errors.New("unauthorized")
	}

	claims := jwt.MapClaims{}
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["expired"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := helper.GenerateToken(&claims)
	if err != nil {
		return response, http.StatusInternalServerError, errors.New("internal server error")
	}

	user.Token = token
	err = service.UserRepository.UpdateToken(user)
	if err != nil {
		return response, http.StatusInternalServerError, errors.New("internal server error")
	}

	response.Token = token

	return response, http.StatusOK, nil
}
