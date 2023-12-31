# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL mantainer="Bel Cattaneo"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...


# Install the package
RUN go install -v ./... 

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# Build the Go app
ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"