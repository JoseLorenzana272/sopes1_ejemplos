# Compilación y generación de Protobuf
FROM golang:1.24-alpine AS builder

# Instalar compilador de Protobuf
RUN apk add --no-cache protoc

# Instalar plugins de Go para gRPC
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="$PATH:/go/bin"
WORKDIR /app

COPY go-server/go.mod ./
COPY go-server/main.go ./
COPY proto/ ./proto/

# Generar archivos .go a partir del .proto internamente en Docker
RUN protoc --go_out=. --go-grpc_out=. proto/report.proto

# Descargar dependencias y compilar binario
RUN go mod tidy
RUN go build -o server_bin main.go

#imagen final
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server_bin .
EXPOSE 50051
CMD ["./server_bin"]