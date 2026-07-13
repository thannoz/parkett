.PHONY: test lint bench up down

test:
	go test -race -count=1 ./...

lint:
	golangci-lint run

bench:
	go test -bench=. -benchmem ./internal/engine/...

up:
	docker compose up -d

down:
	docker compose down
