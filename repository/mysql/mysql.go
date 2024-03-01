package mysql

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_interface "github.com/ekowdd89/go-orm-example/repository/interface"
	"github.com/ekowdd89/go-orm-example/repository/model"
)

type mySQL struct {
	db *gorm.DB
}

func NewDBreadWriter(host, port, user, pass, dbname string) (_interface.UserRepository, error) {
	dnsInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dnsInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &mySQL{
		db: db,
	}, nil
}

func (m *mySQL) FindAll() (model.GetResponse, error) {
	var users []model.User
	result := m.db.Find(&users)
	if result.Error != nil {
		return model.GetResponse{}, result.Error
	}
	rows := model.GetResponse{
		Message: "Success",
		Code:    http.StatusOK,
		Data:    users,
	}
	return rows, nil
}

func (m *mySQL) FindByID(id int) (model.User, error) {
	var user model.User
	result := m.db.First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (m *mySQL) Save(user model.CreateUserRequest) (resp model.UserResponse, err error) {
	result := m.db.Create(&user)
	if result.Error != nil {
		return model.UserResponse{}, result.Error
	}
	resp = model.UserResponse{
		Data: model.User{
			Name:        user.Name,
			Email:       user.Email,
			DisplayName: user.DisplayName,
			Telp:        user.Telp,
		},
	}
	return resp, nil
}

func (m *mySQL) Update(user model.UpdateUserRequest) (model.UserResponse, error) {
	result := m.db.Save(&user)
	if result.Error != nil {
		return model.UserResponse{}, result.Error
	}

	userResponse := model.UserResponse{
		Data: model.User{
			Name:        user.Name,
			Email:       user.Email,
			DisplayName: user.DisplayName,
			Telp:        user.Telp,
		},
	}
	return userResponse, nil
}

func (m *mySQL) Delete(user model.User, id int) (model.DeletedResponse, error) {
	result := m.db.Delete(&user, id)
	if result.Error != nil {
		return model.DeletedResponse{
			Message: fmt.Sprintf("Delete Failed %s", id),
			Code:    http.StatusNotFound,
			Error:   result.Error,
		}, result.Error
	}

	return model.DeletedResponse{
		Message: fmt.Sprintf("Delete Success %v", id),
		Code:    http.StatusOK,
	}, nil
}

func (m *mySQL) Close() error {
	d, _ := m.db.DB()
	return d.Close()
}
