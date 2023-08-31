build:
	docker-compose build main

run:
	docker-compose up main

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5438/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go