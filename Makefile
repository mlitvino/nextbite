backend:
	cd backend && go build -o ./bin/nextbite ./cmd/nextbite

run:
	cd backend && go run ./cmd/nextbite

test-backend:
	cd backend && go test ./...

tidy:
	cd backend && go mod tidy

.PHONY: backend run test-backend tidy-backend
