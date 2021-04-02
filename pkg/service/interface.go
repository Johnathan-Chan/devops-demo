package service

import (
	"github.com/Johnathan-Chan/devops-demo/pkg/core"
	"github.com/Johnathan-Chan/devops-demo/pkg/store"
)

type IService interface {
	Create(core.IObject) error
	Update(core.IObject, core.IObject) error
	Delete(string) error
	Query(store.Condititon) ([]core.IObject, error)
	Range(store.Condititon, func(core.IObject) error) error
}

type Service struct {
	store.IStore
}

func NewService(is store.IStore) *Service {
	return &Service{is}
}
