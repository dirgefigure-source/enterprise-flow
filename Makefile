.PHONY: dev
dev:
	cd backend && go run ./cmd/server

.PHONY: test
test:
	cd backend && go test ./...

.PHONY: tidy
tidy:
	cd backend && go mod tidy

.PHONY: infra-up
infra-up:
	docker compose up -d

.PHONY: infra-down
infra-down:
	docker compose down

.PHONY: infra-logs
infra-logs:
	docker compose logs -f

.PHONY: ps
ps:
	docker compose ps