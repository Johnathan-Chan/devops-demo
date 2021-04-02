package tests

import (
	"github.com/Johnathan-Chan/devops-demo/pkg/api/user"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestStruct(t *testing.T) {
	r := gin.Default()
	g := r.Group("/user")
	{
		g.GET("/:name", user.GetOne)
	}
}
