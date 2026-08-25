.PHONY: run test build benchmark vet

run:
	go run ./cmd/slowapi

test:
	go test ./...

build:
	go build -o bin/slowapi ./cmd/slowapi

benchmark:
	go test ./internal/api -run '^$$' -bench BenchmarkListProducts -benchmem

vet:
	go vet ./...
