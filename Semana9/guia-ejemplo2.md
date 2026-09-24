# Ejemplo - Semana 9 - gRPC en Kubernetes

---

### 1. Objetivo

Demostrar la comunicación interna entre microservicios implementados en Go dentro de un clúster de Kubernetes (GKE), utilizando el protocolo binario de alto rendimiento **gRPC** sobre HTTP/2 y definiciones de mensajes mediante **Protocol Buffers (Protobuf)**.

### 2. Resultado esperado

Tener dos despliegues (Deployments) en ejecución dentro del clúster de Kubernetes:

- **Go Client (Pod):** Genera automáticamente un reporte militar simulado cada 3 segundos y lo envía al servidor vía gRPC.
- **Go Server (Pod):** Recibe la petición binaria a través de un servicio interno (`ClusterIP`), procesa los datos, imprime el contenido en consola y responde con un código de estado (`Status: "1"`).

Toda la comunicación ocurre de manera interna y segura dentro del clúster, sin exponer el servicio gRPC a la red pública.

### 3. Arquitectura del ejemplo

`Go Client (Pod)` -> `gRPC (HTTP/2)` -> `Service ClusterIP (go-server-svc:50051)` -> `Go Server (Pod)`

- **Nombre DNS interno:** `go-server-svc.mumnk8s.svc.cluster.local:50051`
- **Protocolo:** gRPC (RPC binario sobre TCP / HTTP/2)

### 4. Tecnologías

- **Go (Golang 1.24):** Lenguaje utilizado para el cliente y el servidor.
- **gRPC y Protobuf (proto3):** Serialización binaria e invocación remota de procedimientos.
- **Docker Multi-stage:** Compilación de Protobuf y generación de binarios sobre imágenes base Alpine ligeras.
- **Kubernetes (GKE):** Orquestador para ejecutar y enlazar los pods mediante Services y DNS interno.

### 5. Requisitos previos

- Repositorio clonado en tu máquina local.
- CLI de Google Cloud (`gcloud`) instalada y con sesión iniciada (`gcloud auth login`).
- Herramienta `kubectl` y plugin `gke-gcloud-auth-plugin` instalados.
- `docker` instalado (únicamente necesario si deseas compilar y subir tus propias imágenes a Docker Hub).

### 6. Descripción de los componentes

- **`proto/report.proto`:** Define el servicio `WarReportService` con el método RPC `SendReport`, el mensaje de solicitud `WarReportRequest` (país, aviones, barcos, timestamp) y la respuesta `WarReportResponse`.
- **`go-server/`:** Código fuente del servidor gRPC.
  - **`main.go`:** Abre el puerto `50051`, registra el servicio gRPC implementado y atiende las llamadas entrantes devolviendo un acuse de recibo.
  - **`Dockerfile`:** Instala `protoc` y dependencias, compila el archivo `.proto` en Go y genera el binario ligero en Alpine.
- **`go-client/`:** Código fuente del cliente gRPC.
  - **`main.go`:** Resuelve la dirección DNS interna `go-server-svc.mumnk8s.svc.cluster.local:50051`, establece la conexión mediante gRPC y transmite datos en un bucle continuo cada 3 segundos.
  - **`Dockerfile`:** Compila el archivo `.proto` y genera el binario ligero del cliente sobre Alpine.
- **`k8s/` (Manifiestos de Kubernetes):** Carpeta que contiene las definiciones de infraestructura para desplegar y comunicar los microservicios en el clúster:
  - **`00-namespace.yaml`:** Crea el namespace `mumnk8s` para agrupar y aislar los recursos del ejemplo dentro del clúster.
  - **`01-server-deploy.yaml`:** `Deployment` que administra el ciclo de vida del pod del servidor Go (`go-server`) con 1 réplica, exponiendo el puerto `50051`.
  - **`02-server-svc.yaml`:** `Service` de tipo `ClusterIP` con el nombre `go-server-svc`. Proporciona una IP virtual estable y la resolución DNS `go-server-svc.mumnk8s.svc.cluster.local:50051` para enrutar el tráfico gRPC hacia el pod del servidor.
  - **`03-client-deploy.yaml`:** `Deployment` que administra el pod del cliente Go (`go-client`), encargado de emitir continuamente los reportes hacia el servidor.

---

### 7. Guía de ejecución paso a paso

Dado que ya tienes el repositorio clonado, sigue estos pasos desde tu terminal.

#### Paso 1: Ubicarse en el directorio del ejemplo

Abre una terminal y colócate en la carpeta correspondiente a este ejemplo dentro del repositorio:

```bash
cd Semana9
```

Verifica que los archivos estén presentes:

```bash
ls -la
```

---

#### Paso 2: Creación y conexión al clúster en GKE

Si aún no cuentas con un clúster activo en Google Kubernetes Engine (GKE), ejecuta los siguientes comandos:

```bash
# Configurar zona y region por defecto

# Iniciar sesion en Google Cloud (abrira el navegador para autenticarse)
gcloud auth login

# Listar los proyectos disponibles para obtener el ID
gcloud projects list

# Establecer el proyecto activo (reemplaza TU_PROJECT_ID con el ID de tu proyecto)
gcloud config set project TU_PROJECT_ID

# Instalar el plugin de autenticacion de Kubernetes para gcloud
sudo apt-get install google-cloud-cli-gke-gcloud-auth-plugin

# Habilitar las APIs necesarias en Google Cloud (previo a crear el cluster)
gcloud services enable container.googleapis.com compute.googleapis.com logging.googleapis.com

# 1. Crear el cluster en GKE
gcloud container clusters create mumnk8s-cluster \
    --zone us-central1-a \
    --num-nodes 3 \
    --machine-type e2-standard-2 \
    --disk-size=50

# 2. Conectar kubectl al nuevo cluster (obtener credenciales)
gcloud container clusters get-credentials mumnk8s-cluster --zone us-central1-a

# 3. Verificar que los nodos esten listos y conectados
kubectl get nodes
```

---

#### Paso 3: Elección del método de despliegue

Tienes dos opciones para poner en marcha el ejemplo:

##### Opción A: Despliegue directo (Recomendado para pruebas rápidas)

Los manifiestos YAML incluidos en la carpeta `k8s/` ya apuntan a imágenes preconstruidas en Docker Hub (`joselorenzana272/go-server:v2` y `joselorenzana272/go-client:v2`). No es necesario compilar nada localmente si deseas probar el funcionamiento de inmediato.

Pasa directamente al **Paso 4**.

##### Opción B: Construir y subir tus propias imágenes

Si modificas el código en Go o la definición del `.proto`, debes compilar y publicar tus imágenes en tu cuenta de Docker Hub:

1. Iniciar sesión en Docker Hub:

   ```bash
   docker login
   ```

2. Construir la imagen del Servidor:
   El contexto de construcción debe ser el directorio raíz del ejemplo (`Semana9`), ya que el Dockerfile necesita copiar la carpeta `proto/` y `go-server/`:

   ```bash
   docker build -t TU_USUARIO_DOCKER/go-server:v2 -f go-server/Dockerfile .
   docker push TU_USUARIO_DOCKER/go-server:v2
   ```

3. Construir la imagen del Cliente:
   De igual forma, desde la raíz de `Semana9`:

   ```bash
   docker build -t TU_USUARIO_DOCKER/go-client:v2 -f go-client/Dockerfile .
   docker push TU_USUARIO_DOCKER/go-client:v2
   ```

4. Actualizar la referencia de la imagen en los manifiestos YAML:
   - Edita `k8s/01-server-deploy.yaml` y sustituye la propiedad `image:` por `TU_USUARIO_DOCKER/go-server:v2`.
   - Edita `k8s/03-client-deploy.yaml` y sustituye la propiedad `image:` por `TU_USUARIO_DOCKER/go-client:v2`.

---

#### Paso 4: Aplicar los manifiestos en el clúster de Kubernetes

Aplica los recursos en el clúster. Puedes aplicarlos todos juntos o en orden secuencial:

```bash
# Aplicacion directa de todos los manifiestos en la carpeta k8s
kubectl apply -f k8s/
```

---

#### Paso 5: Verificar el estado de los recursos

Verifica que los pods y el servicio se encuentren en ejecución dentro del namespace `mumnk8s`:

```bash
# Consultar estado de los pods
kubectl get pods -n mumnk8s

# Consultar servicios activos
kubectl get svc -n mumnk8s
```

Ambos pods (`go-server-deploy-...` y `go-client-deploy-...`) deben reportar el estado `Running`.

---

#### Paso 6: Monitorear la comunicación gRPC en vivo

Abre dos terminales para observar la interacción entre cliente y servidor en tiempo real:

**Terminal 1 (Logs del Cliente - Emisión de reportes):**

```bash
kubectl logs -f -l app=go-client -n mumnk8s
```

_Salida esperada en el cliente:_

```text
Cliente Go conectado a gRPC. Iniciando envío automático...
Reporte enviado -> País: usa | Aviones: 12 | Barcos: 4 | Status: 1
Reporte enviado -> País: rus | Aviones: 35 | Barcos: 18 | Status: 1
```

**Terminal 2 (Logs del Servidor - Recepción y respuesta):**

```bash
kubectl logs -f -l app=go-server -n mumnk8s
```

_Salida esperada en el servidor:_

```text
Servidor gRPC en Go escuchando en puerto 50051...
Servidor recibió reporte: País=usa, Aviones=12, Barcos=4
Servidor recibió reporte: País=rus, Aviones=35, Barcos=18
```

---

#### Paso 7: Limpieza y eliminación de recursos

Al terminar las pruebas del ejemplo, se recomienda limpiar los recursos para evitar consumos innecesarios en la nube:

##### 1. Eliminar los recursos de Kubernetes dentro del clúster:

```bash
# Eliminar todos los recursos definidos en los manifiestos
kubectl delete -f k8s/

# O eliminar directamente el namespace con todos sus componentes
kubectl delete namespace mumnk8s
```

##### 2. Eliminar el clúster de GKE (para evitar costos en Google Cloud):

```bash
# Eliminar el cluster de GKE por completo
gcloud container clusters delete mumnk8s-cluster --zone us-central1-a --quiet
```
