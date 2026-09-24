FROM golang:1.24-alpine AS builder
RUN apk add --no-cache protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
ENV PATH="$PATH:/go/bin"
WORKDIR /app

COPY go-client/go.mod ./
COPY go-client/main.go ./
COPY proto/ ./proto/

RUN protoc --go_out=. --go-grpc_out=. proto/report.proto
RUN go mod tidy
RUN go build -o client_bin main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/client_bin .
CMD ["./client_bin"]