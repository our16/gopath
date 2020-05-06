package model

type UserVo struct {
	Name string
}

func NewUserVo(name string) *UserVo {
	return &UserVo{
		Name:name,
	}
}
