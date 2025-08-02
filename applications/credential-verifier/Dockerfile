# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /build
COPY ./applications/credential-verifier ./applications/credential-verifier 
COPY ./applications/util ./applications/util 

WORKDIR /build/applications/credential-verifier
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main ./cmd


FROM alpine:latest
WORKDIR /app
COPY --from=builder /main .
EXPOSE 8080
CMD ["./main"]