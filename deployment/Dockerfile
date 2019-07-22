# Builder
FROM golang:1.12 AS builder

WORKDIR /www/go-service

COPY . .
RUN go mod download
RUN go build -o cmd/gin/gin cmd/gin/main.go

# Container
FROM scratch

COPY --from=builder /www/go-service/cmd/gin/gin /www/go-service/cmd/gin/gin

ENTRYPOINT ["/www/go-service/cmd/gin/gin"]