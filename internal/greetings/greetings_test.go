package greetings

import (
	"testing"
	"app/internal/tests"
	"app/internal/config"
)

func init() {
	tests.Setup()
}

func TestHello(t *testing.T) {
	result := Hello("Bla")
	expected := "Hello, Bla!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestHelloTester(t *testing.T) {
	result := Hello(config.Get("USERNAME"))
	expected := "Hello, tester!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
