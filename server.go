package main

import "fmt"

type ServerConfig struct {
	Port int
}

type Server struct {
	config *ServerConfig
}

func DefaultConfig() *ServerConfig {

	c := &ServerConfig{}

	return c
}

func (s *Server) Start() {
	fmt.Println(s.config.Port)
}
