FROM golang:1.18-buster AS builder

ARG VERSION=dev

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main -ldflags=-X=main.version=${VERSION} ./cmd/api/

FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage into this new image
COPY --from=builder /app/main .

# Run the binary when the container starts
CMD ["main"]