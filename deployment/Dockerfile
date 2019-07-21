FROM golang:1.12

WORKDIR /www/go-service
COPY . .

RUN go mod download
RUN go build -o cmd/gin/gin cmd/gin/main.go

CMD ["./cmd/gin/gin"]