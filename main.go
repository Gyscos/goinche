package main

import "flag"

func fillConfig(c *ServerConfig) {
	flag.IntVar(&c.Port, "port", 8080, "The server port")
	flag.Parse()
}

func main() {

	c := DefaultConfig()
	fillConfig(c)

	server := Server{config: c}

	server.Start()
}
