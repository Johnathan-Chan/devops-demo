package authAndUser

import (
	"devops-demo/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	*AuthUser
}

func NewAuth(i api.INormalServer) *Auth {
	g := i.(*AuthUser)
	u := &Auth{g}
	gp := u.Group("/")
	{
		gp.GET("/auths", u.GetAll)
		gp.GET("/auth/:name", u.GetOne)
	}

	return u
}

func (u *Auth) GetAll(g *gin.Context)  {

	g.JSON(http.StatusOK, gin.H{
		"data": "GetAll",
	})
}

func (u *Auth) GetOne(g *gin.Context)  {
	g.JSON(http.StatusOK, gin.H{
		"data": "GetOne",
	})
}
