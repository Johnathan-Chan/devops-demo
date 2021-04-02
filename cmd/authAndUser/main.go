package main

import (
	"devops-demo/pkg/api"
	"devops-demo/pkg/api/authAndUser"
)

func main() {
	//server := api.NewServer()
	//server.Register(&auth.Auth{}, "0.0.0.0:8080")
	//server.Register(&user.User{}, "0.0.0.0:8081")
	//server.Run()
	all := api.NewServer()
	server := authAndUser.NewAuthUser()
	user := authAndUser.NewUser(server)
	auth := authAndUser.NewAuth(server)
	all.Register(user, "0.0.0.0: 8080")
	all.Register(auth, "0.0.0.0: 8081")
	all.Run()
}
