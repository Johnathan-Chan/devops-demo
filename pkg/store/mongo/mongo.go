package mongo

import (
	"github.com/Johnathan-Chan/devops-demo/pkg/core"
	"github.com/Johnathan-Chan/devops-demo/pkg/store"
)

type Mongo struct {
}

func (m Mongo) Apply(database, table string, obj core.IObject) error {
	panic("implement me")
}

func (m Mongo) Get(database, table string, c store.Condititon) (core.IObject, error) {
	panic("implement me")
}

func (m Mongo) Del(database, table string, c store.Condititon) error {
	panic("implement me")
}

func NewMongo() store.IStore {
	return &Mongo{}
}
