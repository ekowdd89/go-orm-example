package middleware

import (
	"fmt"
	"io"
	"net/http"

	_interface "github.com/ekowdd89/go-orm-example/repository/interface"
	"github.com/ekowdd89/go-orm-example/repository/model"
)

type crudRule struct {
	next _interface.UserRepository
	io.Closer
}

type CrudMDRule func(_interface.UserRepository) _interface.UserRepository

func NewCrudMDRule() CrudMDRule {
	return func(ur _interface.UserRepository) _interface.UserRepository {
		return &crudRule{next: ur}
	}
}

func (c *crudRule) Close() error {
	return c.next.Close()
}
func (c *crudRule) FindAll() (model.GetResponse, error) {
	return c.next.FindAll()
}
func (c *crudRule) FindByID(id int) (model.User, error) {
	return c.next.FindByID(id)
}
func (c *crudRule) Save(user model.CreateUserRequest) (model.UserResponse, error) {
	return c.next.Save(user)
}
func (c *crudRule) Update(user model.UpdateUserRequest) (model.UserResponse, error) {
	return c.next.Update(user)
}
func (c *crudRule) Delete(user model.User, id int) (resp model.DeletedResponse, err error) {
	resp, err = c.next.Delete(user, id)
	if err != nil {
		return model.DeletedResponse{
			Message: fmt.Sprintf("Delete Failed %v", id),
			Code:    http.StatusNotFound,
			Error:   err,
		}, err
	}
	return resp, nil
}
