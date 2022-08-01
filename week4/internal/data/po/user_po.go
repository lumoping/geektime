package po

import (
	"time"
	"week4/internal/biz/do"
	"week4/internal/pkg"
)

type UserPO struct {
	Id         int64
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}

func (u *UserPO) ToUserDO() *do.UserDO {
	return &do.UserDO{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUserPO(userDO do.UserDO) *UserPO {
	return &UserPO{
		Id:         pkg.GenID(),
		Name:       userDO.Name,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}
