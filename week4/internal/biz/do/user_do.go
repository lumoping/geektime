package do

type UserDO struct {
	Id   int64
	Name string
}

func (u *UserDO) Valid() bool {
	return len(u.Name) > 0
}
