# Test commands
.PHONY: test test-coverage test-coverage-html test-coverage-check

# Run all tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

# Generate HTML coverage report
test-coverage-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Check if coverage meets minimum threshold (60%)
test-coverage-check:
	@go test ./... -coverprofile=coverage.out
	@echo "Checking coverage threshold..."
	@COVERAGE=$$(go tool cover -func=coverage.out | grep total: | awk '{print $$3}' | sed 's/%//'); \
	if [ $$(echo "$$COVERAGE >= 60" | bc -l) -eq 1 ]; then \
		echo "✅ Coverage $$COVERAGE% meets minimum requirement (60%)"; \
	else \
		echo "❌ Coverage $$COVERAGE% is below minimum requirement (60%)"; \
		exit 1; \
	fi

# Clean coverage files
clean-coverage:
	rm -f coverage.out coverage.html