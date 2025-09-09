FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server server.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
