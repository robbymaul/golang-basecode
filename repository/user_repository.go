package repository

import "basecode/model"

type UserRepository interface {
	Save(user model.User) error
	FindOneByUsername(username string) (model.User, error)
	UpdateToken(user model.User) error
}
