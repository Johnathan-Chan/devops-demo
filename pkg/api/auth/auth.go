package auth

import (
	"devops-demo/pkg/api"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	*gin.Engine
}

func NewAuthServer() api.INormalServer{
	router := gin.Default()
	return &Auth{router}
}

func (a *Auth) Run(addr string) error  {
	return a.Engine.Run(addr)
}



