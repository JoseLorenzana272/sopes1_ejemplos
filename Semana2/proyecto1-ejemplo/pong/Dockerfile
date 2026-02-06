# Imagen base con Go instalado (se usa SOLO para compilar)
# Alpine = ligera
# "AS builder" nombra esta etapa para usarla después
FROM golang:1.25-alpine AS builder

# Define el directorio de trabajo dentro del contenedor
# Todo lo que siga se ejecuta desde /app
WORKDIR /app

# Copia todo el contenido del proyecto (main.go, go.mod, etc.)
# desde el host hacia /app dentro del contenedor
COPY . .

# Compila la aplicación Go
# -o main: el binario final se llamará "main"
# . : compila el módulo/proyecto actual
RUN go build -o main .


# --- Segunda etapa (runtime) ---

# Imagen mínima SOLO para ejecutar el binario
# No tiene Go ni herramientas de compilación
FROM alpine:latest

# Directorio donde vivirá el binario dentro del contenedor
WORKDIR /root/

# Copia únicamente el binario "main" desde la etapa builder
# No se copia el código fuente ni Go
COPY --from=builder /app/main .

# Comando por defecto al iniciar el contenedor
# Ejecuta la aplicación compilada
CMD ["./main"]
