app: build
	./app

build:
	docker exec dev go build -o app cmd/main.go

test:
	echo "Tests coming soon"

up:
	docker-compose up -d

down:
	docker-compose down
