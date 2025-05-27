# Dockerfile
FROM golang:1.24.1

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Command to run the Go application with the -env=docker flag
CMD ["./main", "-env=docker"]