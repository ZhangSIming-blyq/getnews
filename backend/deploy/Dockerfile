# Backend Dockerfile
FROM golang:1.20

WORKDIR /app

# Set the Go proxy to Alibaba Cloud
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd

RUN mv ./cmd/config.json ./config.json

# Expose port (adjust if necessary)
EXPOSE 8080

# Run the application
CMD ["./main"]