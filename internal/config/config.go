package config

import (
	"fmt"
	"errors"
	"github.com/joho/godotenv"
)

var envs map[string]string

func Load(envFile string) error {
	pairs, err := godotenv.Read(envFile)
	if err != nil {
		return errors.New(fmt.Sprintf("Error loading %s file", envFile))
	}
	overload(pairs)

	pairs, err = godotenv.Read(envFile + ".local")
	if err == nil {
		overload(pairs)
	}
	return err	
}

func Get(name string) string {
	return envs[name]
}

func GetAll() map[string]string {
	return envs
}

func overload(pairs map[string]string) {
	if (envs == nil) {
		envs = make(map[string]string)
	}

	for key, value := range pairs {
		envs[key] = value
	}
}

func Clear() {
	envs = make(map[string]string)
}
