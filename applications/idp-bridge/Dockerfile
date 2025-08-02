# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /build
COPY ./applications/idp-bridge ./applications/idp-bridge 
COPY ./applications/util ./applications/util 

WORKDIR /build/applications/idp-bridge
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main ./cmd


FROM alpine:latest
WORKDIR /app
COPY --from=builder /main .
EXPOSE 8080
CMD ["./main"]