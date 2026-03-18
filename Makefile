.PHONY: generate
generate:
	templ generate
	# sqlc generate

.PHONY: build
build: generate
	go build -o server ./cmd/server

.PHONY: run
run: generate
	go run ./cmd/server/main.go

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database "${DATABASE_URL}" up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "${DATABASE_URL}" down
