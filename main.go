package main

import (
	"log"
	"net/http"

	"github.com/ekowdd89/go-orm-example/infra/routes"
	"github.com/ekowdd89/go-orm-example/infra/server"
	repo "github.com/ekowdd89/go-orm-example/repository"
	"github.com/ekowdd89/go-orm-example/repository/usecase"
	"github.com/gorilla/mux"
)

func main() {

	repo, _ := repo.NewCrudRepository(&repo.RepoConfigs{
		DbConfig: repo.DBConfig{
			Host:   "localhost",
			Port:   "3306",
			User:   "root",
			Pass:   "",
			DbName: "mysite",
		},
	})
	defer func() {
		repo.Close()
		log.Println("db connection closed")
	}()
	userUsecase := usecase.NewUserUsecase(repo.DBreadWriter)
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"OK"}`))
	})
	router.HandleFunc("/users", userUsecase.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userUsecase.FindByID).Methods(http.MethodGet)
	router.HandleFunc("/users", userUsecase.Save).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", userUsecase.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", userUsecase.Delete).Methods(http.MethodDelete)
	if err := server.NewServer(server.ServerConfig{
		Handler: routes.NewRouter()(router),
		Address: ":8080",
	}).ListenAndServe(); err != nil {
		log.Println(err)
	}
}
