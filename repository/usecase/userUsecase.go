package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ekowdd89/go-orm-example/utils"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	repo "github.com/ekowdd89/go-orm-example/repository/interface"
	"github.com/ekowdd89/go-orm-example/repository/model"
)

type UserUsecase interface {
	io.Closer
	FindAll(w http.ResponseWriter, t *http.Request)
	FindByID(w http.ResponseWriter, t *http.Request)
	Save(w http.ResponseWriter, t *http.Request)
	Update(w http.ResponseWriter, t *http.Request)
	Delete(w http.ResponseWriter, t *http.Request)
}
type userUsecase struct {
	repo repo.UserRepository
}

func NewUserUsecase(rp repo.UserRepository) UserUsecase {
	return &userUsecase{
		repo: rp,
	}
}

func (u *userUsecase) Close() error {
	return nil
}
func (u *userUsecase) FindAll(w http.ResponseWriter, t *http.Request) {
	rep, err := u.repo.FindAll()
	if err != nil {
		return
	}
	response, _ := json.Marshal(rep)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(response)
	return
}
func (u *userUsecase) FindByID(w http.ResponseWriter, t *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(t)["id"])
	resp, err := u.repo.FindByID(id)
	if err != nil {
		return
	}
	response := model.UserResponse{
		Message: "Success",
		Code:    http.StatusOK,
		Data:    resp,
	}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(result)
}
func (u *userUsecase) Save(w http.ResponseWriter, t *http.Request) {
	req, err := utils.HandlerRequest[model.CreateUserRequest](w, t)
	if err != nil {
		return
	}
	resp, err := u.repo.Save(req)
	if err != nil {
		return
	}
	response := model.UserResponse{
		Message: "Success",
		Code:    http.StatusCreated,
		Data:    resp.Data,
	}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(result)
}
func (u *userUsecase) Update(w http.ResponseWriter, t *http.Request) {
	req, err := utils.HandlerRequest[model.UpdateUserRequest](w, t)
	if err != nil {
		return
	}
	id := mux.Vars(t)["id"]
	req.Id, _ = strconv.Atoi(id)
	req.EmailVerifiedAt = time.Now()
	req.UpdatedAt = time.Now()
	resp, err := u.repo.Update(req)
	if err != nil {
		return
	}
	response := model.UserResponse{
		Message: "Success",
		Code:    http.StatusOK,
		Data:    resp.Data,
	}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(result)
	// return u.repo.Update()
}
func (u *userUsecase) Delete(w http.ResponseWriter, t *http.Request) {
	var (
		response model.DeletedResponse
	)
	id, _ := strconv.Atoi(mux.Vars(t)["id"])
	_, err := u.repo.Delete(model.User{}, id)
	if err != nil {
		response = model.DeletedResponse{
			Message: fmt.Sprintf("Deleted Success %v", id),
			Code:    http.StatusOK,
		}
		log.Err(err)
	} else {
		response = model.DeletedResponse{
			Message: fmt.Sprintf("Deleted Success %v", id),
			Code:    http.StatusOK,
		}
		log.Info().Any("User Usecase funcName ", "Delete()").Msgf("Deleted Success %v", id)
	}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(result)
	// return u.repo.Delete()
}
