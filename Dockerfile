# Use an official Golang runtime as a parent image
FROM golang:1.23-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download any Go modules
RUN go mod download

# Build the Go app
RUN go build -o main .

# Run the app
CMD ["./main"]
