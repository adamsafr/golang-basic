.PHONY: run build-linux-amd64 build-windows-amd64 build-darwin-amd64 test vet

run:
	go run cmd/main.go

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build cmd/main.go

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build cmd/main.go

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build cmd/main.go

test:
	go test ./...

vet:
	go vet ./...
