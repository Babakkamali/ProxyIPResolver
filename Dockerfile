# Build stage
FROM golang:1.22.0-alpine3.19 AS builder
WORKDIR /app
COPY main.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o proxyipchecker .

# Final stage
FROM alpine:3.19.1
RUN apk --no-cache add curl ca-certificates
WORKDIR /root/
COPY --from=builder /app/proxyipchecker .
EXPOSE 3000
CMD ["./proxyipchecker"]
