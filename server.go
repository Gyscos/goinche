package main

import (
	"fmt"
	"io"
	"os"
)

type ServerConfig struct {
	Port    int
	Writer  io.Writer
	OnClose []func()
}

type Server struct {
	config *ServerConfig
}

func (c *ServerConfig) AddOnClose(f func()) {
	c.OnClose = append(c.OnClose, f)
}

func NewServer(config *ServerConfig) *Server {
	s := &Server{
		config: config,
	}

	return s
}

func DefaultConfig() *ServerConfig {

	c := &ServerConfig{
		Port:   8080,
		Writer: os.Stdout,
	}

	return c
}

func (s *Server) Start() {
	fmt.Fprintf(s.config.Writer, "Listening on port %v\n", s.config.Port)
	// ...

	for _, f := range s.config.OnClose {
		f()
	}
}
