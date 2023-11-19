build:
	go build ./cmd/factorial_calculate/main.go

run:
	go run ./cmd/factorial_calculate/main.go

test:
	go test -v -cover ./...

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

# ==============================================================================
# Docker compose commands

dev:
	echo "Starting docker dev environment"
	docker-compose -f deployments/docker-compose.dev.yml up factorial_calculate --build

prod:
	echo "Starting docker prod environment"
	docker-compose -f deployments/docker-compose.prod.yml up factorial_calculate --build

local:
	echo "Starting docker local environment"
	docker-compose -f deployments/docker-compose.local.yml up factorial_calculate --build