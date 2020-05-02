# Params
BINARY_PATH=bin
BINARY_NAME=hello_go

# Commands
build:
	go build -o $(BINARY_PATH)/${BINARY_NAME} cmd/hello-go/main.go

run:
	go run cmd/hello-go/main.go

lint:
	golint ./...

dev:
	gin -b bin/gin-bin --build cmd/hello-go/ -i run