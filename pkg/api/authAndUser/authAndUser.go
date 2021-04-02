package authAndUser

import (
	"github.com/Johnathan-Chan/devops-demo/pkg/api"
	"github.com/gin-gonic/gin"
)

type AuthUser struct {
	*gin.Engine
}

func (au AuthUser) Run(addr string) error {
	return au.Engine.Run(addr)
}

func NewAuthUser() api.INormalServer {
	router := gin.Default()
	return &AuthUser{router}
}
