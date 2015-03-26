package main

import (
	"io/ioutil"
	"testing"
)

func TestServer(t *testing.T) {
	config := DefaultConfig()
	config.Writer = ioutil.Discard
	s := NewServer(config)

	go func() {
		err := s.Start()
		if err != nil {
			t.Error(err)
		}
	}()

	// Test some stuff?

	// We're done
	s.WaitReady()
	err := s.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
