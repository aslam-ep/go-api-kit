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