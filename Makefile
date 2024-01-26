run: build
	@./bin/api/main.go

build: 
	@go build -o bin/api/main.go cmd/main.go

.PHONY: run build
