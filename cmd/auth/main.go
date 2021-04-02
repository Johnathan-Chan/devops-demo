package main

import (
	"devops-demo/pkg/api/auth"
)

func main() {
	auth := auth.NewAuthServer()
	auth.Run("0.0.0.0:8888")
}
