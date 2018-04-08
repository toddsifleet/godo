DIRS=bin

bootstrap:
	glide install

test:
	go test ./...

build:
	go build -o ./bin/client ./cmd/client/main.go
	go build -o ./bin/server ./cmd/server/main.go

$(shell mkdir -p $(DIRS))
