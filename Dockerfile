# Gunakan base image Go dengan Alpine
FROM golang:1.23-alpine3.20

# Install dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary
RUN go build -o main ./cmd

# Jalankan aplikasi
CMD ["./main"]
