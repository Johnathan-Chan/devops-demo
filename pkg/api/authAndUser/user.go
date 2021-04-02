package authAndUser

import (
	"devops-demo/pkg/api"
	"devops-demo/pkg/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	*AuthUser
	user *user.UserService
}

func NewUser(i api.INormalServer) *User {
	g := i.(*AuthUser)
	u := &User{AuthUser: g}
	gp := u.Group("/")
	{
		gp.GET("/users", u.GetAll)
		gp.GET("/user/:name", u.GetOne)
	}

	return u
}

func (u *User) GetAll(g *gin.Context)  {

	g.JSON(http.StatusOK, gin.H{
		"data": "GetAll",
	})
}

func (u *User) GetOne(g *gin.Context)  {
	g.JSON(http.StatusOK, gin.H{
		"data": "GetOne",
	})
}
