package core

import (
	"encoding/json"
	"reflect"
)

type IObject interface {
	GetKind() string
	Get(string) interface{}
	Set(string, interface{})
}

type Object struct {
}

func (o *Object) GetKind() string {
	return reflect.TypeOf(o).Name()
}

func (o *Object) Get(string) interface{} {
	panic("Object Get Method not implement")
}

func (o *Object) Set(string, interface{}) {
	panic("Object Set Method not implement")
}

func Obj2Json(o *Object) ([]byte, error) {
	return json.Marshal(o)
}

func Json2Obj(d []byte, o *Object) error {
	return json.Unmarshal(d, o)
}


