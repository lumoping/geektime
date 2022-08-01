package mysql

import (
	"database/sql"
	"week4/internal/data/po"
)

type Repo struct {
	DB              *sql.DB
	getUserByIdStmt *sql.Stmt
}

func (r Repo) GetUser(id int64) (po.UserPO, error) {
	panic("implement me")
}

func (r Repo) GetUsers() (users []po.UserPO, err error) {
	panic("implement me")
}

func (r Repo) SaveUser(user po.UserPO) error {
	panic("implement me")
}
