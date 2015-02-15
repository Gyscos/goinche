package main

import (
	"flag"
	"log"
	"os"
)

func fillConfig(c *ServerConfig) {
	var output string

	flag.StringVar(&output, "o", "", "File to output to. Leave blank to output to stdout")
	flag.IntVar(&c.Port, "port", 8080, "The server port")
	flag.Parse()

	if output != "" {
		f, err := os.Create(output)
		if err != nil {
			log.Println("Error creating log file:", err)
			log.Println("Fallback to console output.")
		} else {
			c.Writer = f
			c.AddOnClose(func() { f.Close() })
		}
	}
}

func main() {

	c := DefaultConfig()
	fillConfig(c)

	server := NewServer(c)

	server.Start()
}
