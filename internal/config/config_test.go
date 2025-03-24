package config

import (
	"testing"
	"fmt"
)

func initTest() {
	err := Load("../../.env")
	if (err != nil) {
		fmt.Println("Error:", err)
		return
	}
    Load("../../.env.test")
}

func TestEnv(t *testing.T) {
	initTest()
	defer Clear()

	result := Get("ENV")
	expected := "test"
	if result != expected {
		t.Errorf("Expected environment %s, got %s", expected, result)
	}
}

func TestUsername(t *testing.T) {
	initTest()
	defer Clear()

	result := Get("USERNAME")
	expected := "tester"
	if result != expected {
		t.Errorf("Expected username %s, got %s", expected, result)
	}
}
