package redis

import (
	"week4/internal/data/po"
)

type Repo struct {
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
