## hello go app

Simple development environment with dockerized go compiler.

### Installation

- Clone repo
- Copy .env.dist file to .env and configure it`s env vars as you need
- Init the go module file (if you need) *Don`t forget to fill [Your App name]*
> docker run --rm -v $(pwd):/app -w /app golang:1.23 go mod init [Your App name]

### Up environment

> make up

### Build binary file

> make build

### Down environment

> make down
