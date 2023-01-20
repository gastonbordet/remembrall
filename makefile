server-run:
	go run ./cmd/server/main.go

test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out