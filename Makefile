build:
	docker compose build
test:
	docker compose run --rm this go test ./...
lint:
	docker compose run --rm this golangci-lint run
shell:
	docker compose run --rm -it this /bin/ash