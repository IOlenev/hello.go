package main

import (
    "fmt"
    "app/internal/greetings"
    "app/internal/config"
)

func main() {
    config.Load(".env")

    fmt.Println(greetings.Hello(config.Get("USERNAME")));
}
