package service

import (
	"devops-demo/pkg/core"
	"devops-demo/pkg/store"
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

func NewService(is store.IStore) *Service{
	return &Service{is}
}
