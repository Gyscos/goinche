package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

type ServerConfig struct {
	Port    int
	Writer  io.Writer
	OnClose []func()
}

type Server struct {
	// General stuff
	config *ServerConfig
	logger *log.Logger

	// Raw network
	ln net.Listener

	// Lifecycle
	done  chan struct{}
	ready chan struct{}

	// Actual games
}

func (c *ServerConfig) AddOnClose(f func()) {
	c.OnClose = append(c.OnClose, f)
}

func NewServer(config *ServerConfig) *Server {
	s := &Server{
		config: config,
		done:   make(chan struct{}, 1),
		ready:  make(chan struct{}, 1),
		logger: log.New(config.Writer, "Server", log.LstdFlags),
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

func (s *Server) Start() error {
	s.logger.Println("Listening on port", s.config.Port)
	// Finally, we're done.
	defer close(s.done)

	var err error

	// Prepare the handlers. Defined in handlers.go
	s.prepareHandlers()

	// Prepare the server
	srv := &http.Server{Addr: fmt.Sprintf("localhost:%v", s.config.Port)}

	// Listen on the given port
	s.ln, err = net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}

	s.ready <- struct{}{}

	// Actually serve incoming connections with the HTTP server
	err = srv.Serve(s.ln)
	if err != nil {
		return err
	}

	for _, f := range s.config.OnClose {
		f()
	}

	return nil
}

func (s *Server) WaitReady() {
	s.ready <- <-s.ready
}

func (s *Server) Stop() error {
	s.logger.Println("Now stopping server.")
	err := s.ln.Close()
	if err != nil {
		return err
	}
	<-s.done
	return nil
}
