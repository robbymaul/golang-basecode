package services

import "basecode/web"

type AuthService interface {
	Register(request *web.RegisterUserRequest) (status int, err error)
	Login(request web.LoginUserRequest) (response web.LoginUserResponse, status int, err error)
}
