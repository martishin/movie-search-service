# Build stage
FROM golang:1.23.3-alpine3.20 AS builder
WORKDIR /app

# Set the default architecture for local builds
ARG TARGETOS=linux
ARG TARGETARCH=amd64

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Ensure binary is built for the specified platform
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o server cmd/api/main.go

# Final stage
FROM alpine:3.20
WORKDIR /app

# Copy built binary
COPY --from=builder /app/server /app/server

# Ensure binary has execution permission
RUN chmod +x /app/server

EXPOSE 8100
CMD ["/app/server"]
