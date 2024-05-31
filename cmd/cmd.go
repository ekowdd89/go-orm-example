package cmd

import (
	"log"
	"net/http"

	"github.com/ekowdd89/go-orm-example/infra/routes"
	"github.com/ekowdd89/go-orm-example/infra/server"
	repo "github.com/ekowdd89/go-orm-example/repository"
	"github.com/ekowdd89/go-orm-example/repository/usecase"
	"github.com/gorilla/mux"
)

type OptFunc func(*Cmd) error
type Cmd struct {
	repo *repo.RepoConfigs
}

func WithRepo(repo *repo.RepoConfigs) OptFunc {
	return func(c *Cmd) error {

		c.repo = repo
		return nil
	}
}
func New(opts ...OptFunc) (*Cmd, error) {
	cmd := &Cmd{
		repo: &repo.RepoConfigs{
			DbConfig: repo.DBConfig{
				Host:   "localhost",
				Port:   "3306",
				User:   "root",
				Pass:   "",
				DbName: "mysite",
			},
		},
	}
	for _, opt := range opts {
		if err := opt(cmd); err != nil {
			return nil, err
		}
	}
	return cmd, nil
}
func (cmd *Cmd) Run() {
	repo, _ := repo.NewCrudRepository(cmd.repo)
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
