# Go build stage
FROM golang:1.24.6-alpine AS builder

WORKDIR /app
run go mod init server_go
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server server.go

# Final stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
