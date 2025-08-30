.PHONY: run lint dev

dev:
	@which air > /dev/null 2>&1 || go install github.com/cosmtrek/air@latest
	air

run: lint
	go run cmd/standard-service/main.go

lint:
	go fmt ./...
	@which golangci-lint > /dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run