# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the local code into the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose a port (if your Go app listens on a specific port)
EXPOSE 80

# Command to run the executable
CMD ["./main"]