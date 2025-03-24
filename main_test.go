package main

import (
	"testing"
	"app/internal/tests"
)

func init() {
	tests.Setup();
}

func TestMain (m *testing.M) {
	defer tests.Shutdown()

	// some test logic
}
