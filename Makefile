backend:
	cd backend && go build -o ./bin/nextbite ./cmd/nextbite

frontend:
	cd frontend && npm install

run-backend:
	cd backend && go run ./cmd/nextbite

run-frontend:
	cd frontend && npm run dev

run:
	$(MAKE) -j2 run-backend run-frontend

test-backend:
	cd backend && go test ./...

tidy:
	cd backend && go mod tidy

.PHONY: backend frontend run run-backend run-frontend test-backend tidy-backend
