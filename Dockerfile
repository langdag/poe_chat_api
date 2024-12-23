# Use the official Golang image to create a build artifact.
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal alpine image for the final stage
FROM alpine:3.14

# Install ca-certificates for SSL
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Run the binary program
CMD ["./main"]