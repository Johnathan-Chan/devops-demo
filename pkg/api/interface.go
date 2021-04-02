package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

// 接口首字母I大写
type IServer interface {
	Run()
	Register(INormalServer, string)
}

type INormalServer interface {
	Run(addr string) error
}

func NewServer() IServer {
	return &Server{}
}

type Server struct {
	servers map[INormalServer]string
	errChan chan error
}

func (s *Server) Run() {
	s.errChan = make(chan error)

	for server, addr := range s.servers {
		_server := server
		_addr := addr
		go func() {
			s.errChan <- _server.Run(_addr)
		}()
	}

	for  {
		select {
		case err := <-s.errChan:
			if err != nil{
				log.Println(err)
			}
		}
	}
}

func (s *Server) Register(i INormalServer, addr string) {
	if s.servers == nil{
		s.servers = make(map[INormalServer]string)
	}
	s.servers[i] = addr
}


type Gin struct {
	*gin.Engine
}

func NewGinServer() INormalServer {
	route := gin.Default()
	return &Gin{route}
}

func (g *Gin) Run(addr string) error  {
	return g.Engine.Run(addr)
}