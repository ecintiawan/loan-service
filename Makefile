setup:
	@./scripts/setup.sh

mod-gen:
	@echo "Running generate go mod..."
	go mod tidy && go mod vendor -v

wire:
	@echo "Running generate wire dependencies..."
	wire gen ./...

test:
	@echo "Running project unit tests..."
	@go test ./... -timeout 30s -cover -count=1 -race

docker-start:
	@echo "Running docker compose up..."
	docker-compose up

docker-stop:
	@echo "Running docker compose down..."
	docker-compose down

docker-remove:
	@echo "Running remove docker compose..."
	docker-compose down -v

build:
	@echo "Running build http binary..."
	go build -v cmd/http/main.go

run:
	@echo "Running http binary..."
	go run cmd/http/main.go