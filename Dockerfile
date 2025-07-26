# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goshorty


# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/goshorty .
COPY static ./static
EXPOSE 8080
CMD ["./goshorty"]
