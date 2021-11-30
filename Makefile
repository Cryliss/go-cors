build:
	go mod download
	go build -o gocors bin/go-cors/main.go
