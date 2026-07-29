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

## Cómo ejecutar localmente

Para ejecutar cualquiera de los servicios de forma local, navega a la carpeta correspondiente y ejecuta:

```bash
go run main.go
# O si el archivo se llama ping.go / pong.go
go run ping.go
```

**Nota:** Asegúrate de configurar correctamente la IP del servicio Pong (`TargetIP`) en el código del servicio Ping para que puedan comunicarse entre sí, especialmente si están en diferentes máquinas o contenedores.

## Cómo instalar containerd en Ubuntu Server

Para instalar `containerd` en un servidor Ubuntu, sigue estos pasos desde la terminal:

1. Actualiza el índice de paquetes del sistema:

   ```bash
   sudo apt-get update
   ```

2. Instala el paquete oficial de `containerd`:

   ```bash
   sudo apt-get install -y containerd
   ```

3. Crea el directorio y el archivo de configuración por defecto:

   ```bash
   sudo mkdir -p /etc/containerd
   containerd config default | sudo tee /etc/containerd/config.toml
   ```

4. Reinicia y habilita el servicio para que inicie automáticamente:

   ```bash
   sudo systemctl restart containerd
   sudo systemctl enable containerd
   ```

5. Verifica que el servicio esté activo y funcionando correctamente:
   ```bash
   sudo systemctl status containerd
   ```
   También puedes revisar la versión instalada ejecutando:
   ```bash
   containerd --version
   ```

## Cómo construir y ejecutar imágenes (Docker + containerd)

En este flujo construiremos las imágenes utilizando **Docker**, las subiremos a Docker Hub y luego las ejecutaremos utilizando **containerd** (mediante su herramienta nativa `ctr`).

### 1. Construir y subir las imágenes con Docker

Navega a la ruta de cada servicio, compila las imágenes con el tag de tu repositorio y súbelas:

```bash
# Para el servicio Pong
cd fiber/pong
docker build -t *usuario*/pong-app:latest .
docker push *usuario*/pong-app:latest

# Para el servicio Ping (asumiendo que regresas a la carpeta del proyecto)
cd ../../fiber/ping
docker build -t *usuario*/ping-app:latest .
docker push *usuario*/ping-app:latest
```

### 2. Cargar las imágenes en containerd (en cada VM correspondientemente)

Dependiendo de en qué máquina virtual (VM) te encuentres, descarga la imagen correspondiente.

#### En VM 1 (Servidor Pong):

```bash
sudo ctr images pull docker.io/*usuario*/pong-app:latest
```

#### En VM 2 (Cliente Ping):

```bash
sudo ctr images pull docker.io/*usuario*/ping-app:latest
```

#### Verificar descarga de imágenes:

Para verificar que las imágenes se descargaron correctamente en la VM local, lista las imágenes disponibles en containerd con:

```bash
sudo ctr images list
```

### 3. Ejecutar los contenedores con containerd

Utiliza `ctr` para correr las imágenes en sus respectivas VMs:

#### En VM 1 (Servidor Pong):

```bash
# Ejecuta el contenedor del servicio Pong
sudo ctr run --net-host docker.io/elianreyes/pong-app:latest pong-service
```

#### En VM 2 (Cliente Ping):

```bash
# Ejecuta el contenedor del servicio Ping
sudo ctr run --net-host docker.io/elianreyes/ping-app:latest ping-service
```

**Nota sobre red:** El uso del flag `--net-host` permite que el contenedor utilice la red principal del servidor directamente. Esto simplifica la configuración para que el servicio Ping pueda alcanzar al servicio Pong utilizando la IP del anfitrión (VM 1).

### 4. Administrar contenedores y tareas en containerd

Para listar, verificar el estado y detener/eliminar los contenedores creados con `ctr` en la VM actual, utiliza los siguientes comandos:

#### Ver contenedores y tareas creadas:

```bash
# Listar los contenedores registrados en containerd
sudo ctr containers list

# Listar las tareas activas (procesos de los contenedores corriendo)
sudo ctr tasks list
```

#### Detener y eliminar contenedores:

Para detener por completo un contenedor, debes detener su tarea asociada y luego eliminar tanto la tarea como el contenedor:

```bash
# 1. Detener la tarea del contenedor (reemplaza <nombre-servicio> por pong-service o ping-service)
sudo ctr tasks kill <nombre-servicio>

# 2. Eliminar la tarea
sudo ctr tasks rm <nombre-servicio>

# 3. Eliminar el contenedor
sudo ctr containers rm <nombre-servicio>
```
