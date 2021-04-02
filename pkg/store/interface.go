package store

import "devops-demo/pkg/core"

type Condititon = map[string]string

type IStore interface {
	Apply(database, table string, obj core.IObject) error
	Get(database, table string, c Condititon) (core.IObject, error)
	Del(database, table string, c Condititon) error
}
