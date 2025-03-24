# hello go app

Simple Go development environment with dockerized go compiler.
The app reads enviroment vars from the env file and greets the user

Unfortunately, the debugger feature does not working ((

## Installation

- Clone repo
- Copy .env file to .env.local and configure it`s env vars as you need

## Up environment

> make up

## Testing code

> make test

## Build binary file (destinaton - project`s root)

> make build

## Down environment

> make down

## Reminder

To add package to project "app" follow next steps:

- create package subfolder in the root of the project "bla"
- create go file ./bla/bla.go
- add package path in the import section of main.go root file
- exec command in app root

    > go mod tidy
