package _interface

import (
	"io"

	"github.com/ekowdd89/go-orm-example/repository/model"
)

type UserRepository interface {
	io.Closer
	FindAll() (model.GetResponse, error)
	FindByID(id int) (model.User, error)
	Save(user model.CreateUserRequest) (model.UserResponse, error)
	Update(user model.UpdateUserRequest) (model.UserResponse, error)
	Delete(user model.User, id int) (model.DeletedResponse, error)
}
