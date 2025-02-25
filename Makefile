build:
	go build -o bin/gim cmd/gim/main.go

run:
	./bin/gim

test:
	go test -v ./...
