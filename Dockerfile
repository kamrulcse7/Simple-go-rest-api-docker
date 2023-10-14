# Use the official Golang image as the base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /simple-go-rest-api-docker

# Copy the local package files to the container's working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]