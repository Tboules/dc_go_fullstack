run: build
	@./bin/api/main.go

watch_tailwind:
	@cmd/web/tailwindcss -i cmd/web/style/input.css -o cmd/web/style/output.css --watch

templates:
	@cmd/web/tailwindcss -i cmd/web/style/input.css -o cmd/web/style/output.css --minify
	@templ generate

build: 
	@go build -o bin/api/main.go cmd/main.go 

buildmigration:
	@go build -o bin/db/migrate.go cmd/db/migrate.go

watch:
	@air

migrate_up: buildmigration
	@./bin/db/migrate.go up

migrate_down: buildmigration
	@./bin/db/migrate.go down

.PHONY: migrate_up migrate_down buildmigration run build watch templates_dev templates

