package user

import (
	"devops-demo/pkg/api"
	"github.com/gin-gonic/gin"
)

type User struct {
	*gin.Engine
}

func NewUserServer() api.INormalServer {
	route := gin.Default()
	InitUserRoute(route)
	return &User{route}
}

func (u *User) Run(addr string) error  {
	return u.Engine.Run(addr)
}


