test:
	go test ./...

gen:
	go generate ./...

dev:
	go run cmd/gin/main.go

build:
	go build -o cmd/gin/gin cmd/gin/main.go

run:
	cmd/gin/gin

bench:
	go test -bench=. ./benchmark
