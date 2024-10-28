# Development for Mac
init:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest
generate:
	swag fmt
	swag init
	sqlc generate
test:
	go test ./internal/domain/...
test_cover:
	go test -v -cover=true ./internal/domain/...
build:
	go build -o main
build_and_run:
	go build -o main
	./main
build_docker_compose:
	docker-compose build --no-cache
run_docker_compose:
	docker-compose up -d
stop_docker_compose:
	docker-compose down
