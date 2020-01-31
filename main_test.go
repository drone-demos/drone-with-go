package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "hello drone world test" {
		t.Errorf("got %s expected %s", HelloWorld(), "hello drone world test") 
	}
}
