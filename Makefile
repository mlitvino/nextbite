backend:
	cd backend && go build -o ./bin/nextbite .

run:
	cd backend && go run .

test-backend:
	cd backend && go test ./...

tidy-backend:
	go mod tidy

.PHONY: backend run test tidy
