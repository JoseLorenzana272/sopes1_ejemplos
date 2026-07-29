# Ejemplo Proyecto 1 - Ping Pong API

Este directorio contiene ejemplos de implementación para un sistema simple de comunicación entre dos servicios (microservicios) escritos en Go, simulando una arquitectura básica de **Ping-Pong**.

El objetivo es demostrar cómo un servicio (`ping`) realiza una petición HTTP a otro servicio (`pong`), y cómo ambos pueden ser implementados utilizando diferentes enfoques en Go.

## Estructura de Carpetas

El proyecto está dividido en dos implementaciones principales según la librería utilizada para el servidor web:

```
proyecto1-ejemplo/
├── fiber/          # Implementación usando el framework Fiber
│   ├── ping/       # Servicio Cliente (Ping) con Fiber
│   └── pong/       # Servicio Servidor (Pong) con Fiber
└── http-net/       # Implementación usando la librería estándar net/http
    ├── ping/       # Servicio Cliente (Ping) estándar
    └── pong/       # Servicio Servidor (Pong) estándar
```

## Descripción de los Servicios

### 1. Servicio Ping (Cliente)
Este servicio actúa como el iniciador de la comunicación.
- **Función**: Recibe una petición del usuario y llama internamente al servicio Pong.
- **Puerto**: `8081`

**Endpoints:**
- `GET /iniciar`: Endpoint principal.
    1. Recibe la petición.
    2. Realiza una llamada HTTP GET al servicio Pong (`http://<IP_PONG>:8082/responder`).
    3. Retorna al usuario la respuesta combinada ("Ping: Llame a la otra API y me dijo: [Respuesta de Pong]").
- `GET /health` (Solo versión Fiber): Verifica que el servicio esté activo (`UP`).

### 2. Servicio Pong (Servidor)
Este servicio actúa como el receptor.
- **Función**: Espera llamadas y responde con un mensaje simple.
- **Puerto**: `8082`

**Endpoints:**
- `GET /responder`: Retorna el mensaje "¡Pong! (Desde Containerd en VM 2)".
- `GET /health` (Solo versión Fiber): Verifica que el servicio esté activo (`UP`).

## Implementaciones

### Versión Fiber (`/fiber`)
Utiliza el framework [Go Fiber](https://gofiber.io/), que es un framework web inspirado en Express.js, conocido por su alto rendimiento y facilidad de uso.
- Incluye middleware de `logger` para registrar peticiones.
- Utiliza `fiber.Ctx` para el manejo de contexto HTTP.
- Implementa endpoints de salud (`/health`) devolviendo JSON.

### Versión Standard Library (`/http-net`)
Utiliza el paquete nativo `net/http` de Go.
- Es la forma más básica y nativa de crear servidores en Go.
- Utiliza `http.HandleFunc` y `http.ResponseWriter`.
- Ideal para entender los fundamentos sin dependencias externas.

## Cómo ejecutar

Para ejecutar cualquiera de los servicios, navega a la carpeta correspondiente y corre:

```bash
go run main.go
# O si el archivo se llama ping.go / pong.go
go run ping.go
```

**Nota:** Asegúrate de configurar correctamente la IP del servicio Pong (`TargetIP`) en el código del servicio Ping para que puedan comunicarse entre sí, especialmente si están en diferentes máquinas o contenedores.