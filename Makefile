run: build
	@./bin/api/main.go

build: 
	@go build -o bin/api/main.go cmd/main.go

watch:
	@air

.PHONY: run build
