run: build
	@./bin/api/main.go

watch_tailwind:
	@cmd/web/tailwindcss -i cmd/web/style/input.css -o cmd/web/style/output.css --watch

templates:
	@cmd/web/tailwindcss -i cmd/web/style/input.css -o cmd/web/style/output.css --minify
	@templ generate

build: 
	@go build -o bin/api/main.go cmd/main.go 

watch:
	@air

# migrateup:
# 	@migrat

.PHONY: migrateup run build watch templates_dev templates

