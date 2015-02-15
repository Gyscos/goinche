package main

import (
	"io/ioutil"
	"testing"
)

func TestServer(t *testing.T) {
	config := DefaultConfig()
	config.Writer = ioutil.Discard
	s := NewServer(config)
	s.Start()
}
