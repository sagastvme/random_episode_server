# Use the official Golang image as the build stage
FROM golang:1.24 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker's layer caching for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o /app/main main.go

# Use a minimal base image for the final runtime stage
FROM debian:bookworm-slim

# Set the working directory in the runtime container
WORKDIR /app

# Install any necessary system dependencies (optional)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the built application from the build stage
COPY --from=build /app/main /app/main

# Copy static directories (views and videos)
COPY --from=build /app/views /app/views
COPY --from=build /app/videos /app/videos

# Expose the port the app runs on
EXPOSE 3000

# Define the entry point for the container
CMD ["/app/main"]
