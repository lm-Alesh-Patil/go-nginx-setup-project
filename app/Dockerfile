# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init go-nginx-demo || true
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Set default port (can be overridden by environment variable)
ENV PORT=8080
EXPOSE $PORT
# Start app using the PORT from environment
CMD ["sh", "-c", "./main --port=$PORT"]
