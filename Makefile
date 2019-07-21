test:
	go test ./...

gen:
	go generate ./...

dev:
	go run bin/gin/main.go

build:
	go build -o bin/gin/gin bin/gin/main.go

run:
	bin/gin/gin
