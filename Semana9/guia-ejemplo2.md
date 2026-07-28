# Ejemplo 2 - Semana 9 - gRPC

---

### 1. Objetivo

Implementar la capa intermedia del proyecto (Go Deployment 1 y Go Deployment 2) comunicándolos internamente dentro de GKE mediante el protocolo binario **gRPC**, utilizando exactamente la estructura de Protobuf.

### 2. Resultado

Tener dos pods corriendo en GKE. El "Go Client" generará automáticamente un reporte cada 3 segundos y lo enviará vía gRPC. El "Go Server" recibirá la petición binaria, la imprimirá en consola y le responderá un status. Todo esto sucederá internamente sin salir a internet.

### 3. Arquitectura de este ejemplo

`Go Client (Pod)` → `gRPC (HTTP/2)` → `Service (ClusterIP)` → `Go Server (Pod)`.

### 4. Tecnologías

- **Go (Golang)**
- **gRPC y Protobuf**
- **Docker Multi-stage con Protoc**
- **Kubernetes**

### 5. Requisitos previos

- `kubectl` apuntando al clúster de GKE.
- `docker` logueado para subir imágenes (`docker login`). (Usar las de joselorenzana272)
- namespace `mumnk8s` ya creado.

### 6. Estructura final de carpetas del proyecto

```
demo2-gcp/
├── proto/
│   └── report.proto
├── go-server/
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── go-client/
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
└── k8s/
		|__ 00-namespace.yaml
    ├── 01-server-deploy.yaml
    ├── 02-server-svc.yaml
    └── 03-client-deploy.yaml
```

### 7. Código

- **`proto/report.proto`:** Define el mensaje, el Enum de países y el servicio `WarReportService`.
- **`go-server/main.go`:** Abre el puerto `50051`. Implementa la interfaz gRPC generada y responde con `Status: "1"`.
- **`go-client/main.go`:** Se conecta al DNS interno del clúster (`go-server-svc.mumnk8s.svc.cluster.local:50051`) y envía un struct en un bucle infinito cada 3 segundos.

### 8. Dockerfile

Ambos Dockerfiles instalan `protoc`, generan el código Go a partir del archivo `.proto`, descargan dependencias y compilan el binario final en una imagen ligera de Alpine Linux.

### 9. Manifiestos YAML

- **01-server-deploy:** Pod del servidor Go.
- **02-server-svc:** Service interno `ClusterIP` exponiendo el puerto 50051.
- **03-client-deploy:** Pod del cliente Go.

### 10. Comandos

```bash
# 1. Crear carpetas
mkdir -p demo2-gcp/proto demo2-gcp/go-server demo2-gcp/go-client demo2-gcp/k8s

# 2. Entrar a la raíz de la demo2
cd demo2-gcp

# 3. Construir y subir imagen del Server
docker build -t joselorenzana272/go-server:v1 -f go-server/Dockerfile .
docker push joselorenzana272/go-server:v1

# 4. Construir y subir imagen del Client
docker build -t TU_USUARIO/go-client:v1 -f go-client/Dockerfile .
docker push TU_USUARIO/go-client:v1

# 5. Aplicar YAMLs en GKE
cd k8s
kubectl apply -f .
```

### 11. Orden de ejecución

1. Crear los archivos respetando estrictamente la estructura de carpetas (el Dockerfile necesita ver la carpeta `proto` y la carpeta `go-server` al mismo tiempo).
2. Modificar los YAMLs para poner usuario de Docker Hub. (esto ya está con las ya creadas)
3. Ejecutar los builds y subir las imágenes. (usar las que ya están creadas)
4. Desplegar primero el Server y su Service en GKE.
5. Desplegar el Client en GKE.
6. Revisar los logs para validar la comunicación.

### 12. Comandos para ver la comunicación en vivo

Abrir dos terminales para ver cómo platican entre ellos en vivo:

```bash
# Terminal 1 (Ver qué envía el cliente):
kubectl logs -f -l app=go-client -n mumnk8s

# Terminal 2 (Ver qué recibe el servidor):
kubectl logs -f -l app=go-server -n mumnk8s
```

### 13. Qué salida o comportamiento debería ver si todo va bien

- **En el Cliente se verá:** `Reporte enviado a GKE. Status servidor: 1`
- **En el Servidor se verá:** `Servidor recibió reporte: País=usa, Aviones=12, Barcos=4`