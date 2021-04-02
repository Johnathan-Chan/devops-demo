package user

import (
	"devops-demo/pkg/core"
	"devops-demo/pkg/service"
	"devops-demo/pkg/store"
)

type UserService struct {
	service.Service
}

func (u UserService) Create(object core.IObject) error {
	panic("implement me")
}

func (u UserService) Update(object core.IObject, object2 core.IObject) error {
	panic("implement me")
}

func (u UserService) Delete(s string) error {
	panic("implement me")
}

func (u UserService) Query(condititon store.Condititon) ([]core.IObject, error) {
	panic("implement me")
}

func (u UserService) Range(condititon store.Condititon, f func(core.IObject) error) error {
	panic("implement me")
}

func NewUserService(s service.Service) service.IService {
	return &UserService{s}
}


