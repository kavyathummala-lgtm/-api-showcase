FROM golang:1.23-alpine AS builder
RUN apk add --no-cache protobuf protobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product.proto
RUN go build -o server ./server

FROM alpine:3.20
RUN apk add --no-cache netcat-openbsd
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 50052
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
  CMD nc -z localhost 50052 || exit 1
CMD ["./server"]
