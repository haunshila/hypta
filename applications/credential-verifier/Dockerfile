# Stage 1: Build
FROM golang:1.23.1 AS builder
WORKDIR /app
# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy only Go-related files, excluding frontend
COPY applications/credential-issuance/ applications/credential-issuance/
COPY applications/util/ applications/util/
# Build the credential-issuance binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/credential-issuance ./applications/credential-issuance

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app
# Copy the built binary from the builder stage
COPY --from=builder /app/credential-issuance .
# Expose the port the API Gateway listens on
EXPOSE 8080
# Run the credential-issuance binary
CMD ["./credential-issuance"]