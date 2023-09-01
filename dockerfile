# Build
FROM golang:1.21-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go build -ldflags "-s -w" -o main cmd/app/main.go

# Package
FROM alpine:3
COPY --from=builder /build/main /app/main
COPY --from=builder /build/migrations /app/migrations
WORKDIR /app
CMD ["./main"]
