package repository

import "github.com/ropel12/Latihan/entity"

type UserInterface interface {
	FindByUsername(username string) (res entity.User, err error)
	CreateUser(data entity.User) error
	UpdateUser(data entity.User) error
}
