package main

import (
	"github.com/Johnathan-Chan/devops-demo/pkg/api/user"
)

func main() {
	user := user.NewUserServer()
	user.Run("0.0.0.0:8081")
}
