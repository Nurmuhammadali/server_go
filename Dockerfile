# Build stage
FROM golang:1.24.6-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod tidy       # go.sum shu yerda hosil bo‘ladi

COPY . .
RUN go build -o server server.go

# Final stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
