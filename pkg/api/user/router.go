package user

import "github.com/gin-gonic/gin"

func InitUserRoute(r *gin.Engine){
	u := r.Group("/user")
	{
		u.GET("/:name", GetOne)
	}
}
