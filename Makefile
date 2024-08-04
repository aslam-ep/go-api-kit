# up: starts all container in the background without forcing build
up:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "Docker containers started!"

# down: stop docker compose
down:
	@echo "Stopping docker containers..."
	docker-compose down
	@echo "Docker containers stopped!"

# up_build: stop docker compose if exists, builds and start docker compose
up_build:
	@echo "Stopping docker containers(if running...)"
	docker-compose down
	@echo "Building and starting docker containers..."
	docker-compose up --build -d
	@echo "Docker containers started!"

# swagger: installing swagger for api documentation
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

# swagger_init: initiating swagger docs
swagger_init:
	swag init -g ./cmd/main.go -o ./docs/swagger

# migrate_up: to run the migrations up command
migrate_up:
	migrate -path database/migrations -database "postgresql://root:password@localhost:5432/e-commerce?sslmode=disable" -verbose up

# migrate_down: to run the migrations down command
migrate_down:
	migrate -path database/migrations -database "postgresql://root:password@localhost:5432/e-commerce?sslmode=disable" -verbose down
