
PKG := $(shell go list ./...)
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

## ----- GOLANG
.PHONY: start test coverage coverage-html clear

start: ## Run the server
	go run ./cmd/webserver/main.go

test: ## Run the tests
	go test -v -cover -coverprofile=$(COVERAGE_FILE) $(PKG)

coverage: ## Run the coverage report
	go tool cover -func=$(COVERAGE_FILE)

coverage-html: test ## Generate HTML coverage report and open in browser
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	open $(COVERAGE_HTML)

clear: ## Clear up coverage files
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

## ----- DOCKER
.PHONY: build status logs run up down stop clean
build: ## Build docker image
	docker compose build

status: ## Get status of containers
	docker compose ps

logs: ## Get logs of containers
	docker compose logs --follow

run: ## Run the docker image locally
	docker run -p 8080:8080 --env-file .env weather-cloud-run:latest

up: build ## Build and start docker containers
	docker compose up -d

down: ## Put the compose containers down
	docker compose down

stop: ## Stop docker containers
	docker compose stop

clean: stop ## Stop docker containers, clean data and workspace
	docker compose down -v --remove-orphans

## ----- MOCKERY
MOCKERY_VERSION := v2.53.2
BIN := $(shell go env GOPATH)/bin/mockery

.PHONY: install-mockery reinstall-mockery generate-mocks

install-mockery: ## Install mockery at fixed version
	go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)

reinstall-mockery: ## Force reinstall mockery at fixed version
	rm -f $(BIN)
	go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)

generate-mocks: ## Generate mocks using go:generate
	go generate ./...
