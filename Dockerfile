# Start from golang base image
FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Download all the dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o app ./cmd/server

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT [ "./app", "-config", "api-project-dev"]