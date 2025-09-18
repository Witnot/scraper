# Dockerfile

# Build stage
FROM golang:1.24-bookworm as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/app

# Runtime stage (same Go base to avoid glibc mismatch)
FROM golang:1.24-bookworm
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
