package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOne(g *gin.Context){
	g.JSON(http.StatusOK, gin.H{})
}