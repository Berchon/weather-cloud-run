# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install CA certificates to be copied later
RUN apk add --no-cache ca-certificates

# Download dependencies first (cache layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build static binary for Linux amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/webserver/main.go


# ---------- Final stage ----------
FROM scratch

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/server .

# Copy CA certificates bundle (needed for HTTPS requests)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the default Cloud Run port
EXPOSE 8080

# Use a non-root user to run the application for better security
USER nonroot:nonroot

# Run the binary
ENTRYPOINT ["./server"]
