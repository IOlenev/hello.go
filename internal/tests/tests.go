package tests

import (
	"fmt"
	"app/internal/config"
)

func Setup() {
	// Setting up envs, db connections, etc.
	envFile := "../../.env"
	err := config.Load(envFile)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	envFile = "../../.env.test"
	err = config.Load(envFile)
	if (err != nil) {
		fmt.Println(err)
	}
}

func Shutdown() {
	config.Clear()
}
