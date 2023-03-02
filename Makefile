build:
	go build -o playlist.exe ./cmd/playlist

test:
	go test ./internal/playlist

.DEFAULT_GOAL := build