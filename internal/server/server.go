package server

import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
	Port   string
	Host   string
}

func ServerInit(port, host string) Server {
	s := Server{Port: port, Host: host}
	s.Engine = gin.Default()
	return s
}

func (s *Server) Run() error {
	if err := s.Engine.Run(s.Host + ":" + s.Port); err != nil {
		return err
	}
	return nil
}
