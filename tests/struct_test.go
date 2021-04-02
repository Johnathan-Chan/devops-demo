package tests

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestStruct(t *testing.T) {
	r := gin.Default()
	g := r.Group("/user")
	{
		g.GET("/:name", func(context *gin.Context) {
			context.JSON(200, gin.H{})
		})
	}
}
