# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Get all the dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the GO app
RUN go build -o main .

# Expose... port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

