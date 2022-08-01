package biz

import "week4/internal/biz/do"

type UserBiz struct {
	repo UserData
}

func (b *UserBiz) GetUserById(id uint64) (*do.UserDO, error) {
	return b.repo.GetUser(id)
}

func (b *UserBiz) GetUsers() ([]do.UserDO, error) {
	return b.repo.GetUsers()
}

func NewUserBiz(repo UserData) *UserBiz {
	return &UserBiz{
		repo: repo,
	}
}

type UserData interface {
	GetUser(id uint64) (*do.UserDO, error)
	GetUsers() (users []do.UserDO, err error)
	SaveUser(user do.UserDO) error
}
