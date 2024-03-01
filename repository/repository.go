package repository

import (
	"io"

	_interface "github.com/ekowdd89/go-orm-example/repository/interface"
	"github.com/ekowdd89/go-orm-example/repository/mysql"
)

type Repository struct {
	DBreadWriter _interface.UserRepository
	io.Closer
}

func NewCrudRepository(rp *RepoConfigs) (*Repository, error) {
	const (
		funcName = "NewCrudRepository"
	)

	DbReadWriter, err := mysql.NewDBreadWriter(
		rp.DbConfig.Host,
		rp.DbConfig.Port,
		rp.DbConfig.User,
		rp.DbConfig.Pass,
		rp.DbConfig.DbName,
	)
	if err != nil {
		return nil, err
	}
	return &Repository{
		DBreadWriter: DbReadWriter,
	}, nil
}

func (r *Repository) Close() {
	r.DBreadWriter.Close()
}
