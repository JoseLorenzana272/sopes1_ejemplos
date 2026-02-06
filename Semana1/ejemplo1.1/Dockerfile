# Usamos Alpine (Go) como base
FROM golang:1.25-alpine

# Configuración del entorno
WORKDIR /app
# copiar archivos (fuente) del dir actual
COPY . .

# Instalar dependencias 
RUN go mod download

# Compilar
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Puerto expuesto (debe ser el mismo que el del servidor)
EXPOSE 80

# Comando de ejecución
CMD ["/app/server"]
