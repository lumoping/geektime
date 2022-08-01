package data

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"week4/internal/biz/do"
	"week4/internal/data/mysql"
	"week4/internal/data/po"
	"week4/internal/data/redis"
)

type UserData struct {
	mysql mysql.Repo
	redis redis.Repo
}

func (u *UserData) GetUser(id int64) (*do.UserDO, error) {
	var user po.UserPO
	var err error
	user, err = u.redis.GetUser(id)
	if err != nil {
		user, err = u.mysql.GetUser(id)
	}
	if err != nil {
		return nil, err
	}
	return user.ToUserDO(), nil
}

func (u *UserData) GetUsers() ([]do.UserDO, error) {
	var users []po.UserPO
	var err error
	users, err = u.redis.GetUsers()
	if err != nil {
		users, err = u.mysql.GetUsers()
	}
	if err != nil {
		return nil, err
	}
	usersDO := make([]do.UserDO, len(users))
	for i, user := range users {
		usersDO[i] = *user.ToUserDO()
	}
	return usersDO, nil
}

func (u *UserData) SaveUser(user do.UserDO) error {
	if !user.Valid() {
		return errors.New("name can not be empty")
	}
	return u.mysql.SaveUser(*po.NewUserPO(user))
}

type UserRepo interface {
	GetUser(id int64) (*po.UserPO, error)
	GetUsers() (users []po.UserPO, err error)
	SaveUser(user po.UserPO) error
}
