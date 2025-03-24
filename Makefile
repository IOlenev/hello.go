app: build
	./app

build:
	docker exec compiler go build -o app main.go

test:
	docker exec compiler bash -c 'go clean -testcache && go test ./...'

up:
	docker-compose up -d

down:
	docker-compose down

bash:
	docker exec -it compiler bash

