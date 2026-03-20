# Ejemplos - Guias - SOPES1

---

### 1. Ejemplo 1

Construir la capa de entrada pública en Google Cloud: una API REST en Rust desplegada en GKE, expuesta a internet mediante el Gateway API nativo de GCP, y configurada con HPA para escalar automáticamente ante el tráfico de estrés generado por Locust.

### 2. Resultado esperado

Locust enviará miles de peticiones desde tu computadora local hacia una **IP Pública real de Google Cloud**. El Load Balancer enrutará el tráfico a los pods de Rust en GKE. El consumo de CPU subirá, el HPA escalará la API a 3 réplicas en vivo, y al detener Locust, se reducirá a 1.

### 3. Arquitectura del ejemplo 1

`Locust (Local)` → `GCP External Load Balancer (Gateway API)` → `HTTPRoute` → `Service` → `Rust Pods (GKE)` ↔ `HPA (Monitoreando CPU)`.

### 4. Tecnologías

- **GKE (Google Kubernetes Engine):** Entorno final del proyecto. Instala y gestiona todo automáticamente.
- **Gateway API nativo (gke-l7-global-external-managed):** En GKE ya no hay que instalar controladores extraños (como Envoy); Google traduce el YAML del Gateway directamente en un Balanceador de Carga en la nube.
- **Rust (Axum) y Locust**
- **Docker Hub (Temporal):** Como Zot se implementa en otra fase, hoy subiremos la imagen a Docker Hub público para que GKE pueda descargarla fácilmente.

### 5. Requisitos previos

- `gcloud` CLI: Instalado y autenticado (`gcloud auth login` y `gcloud config set project TU_PROYECTO`).
- `kubectl`: Instalado (`gcloud components install kubectl`).
- `docker`: Instalado y logueado en tu cuenta (`docker login`).
- `rustup` / `cargo`: Para compilar Rust (`cargo --version`).
- `python3` y `pip`: Para Locust (`python3 --version`).

### 6. Estructura de carpetas del ejemplo1

```
demo1-gcp/
├── api-rust/
│   ├── Cargo.toml
│   ├── Dockerfile
│   └── src/
│       └── main.rs
├── k8s/
│   ├── 01-namespace.yaml
│   ├── 02-deployment.yaml
│   ├── 03-service.yaml
│   ├── 04-gateway.yaml
│   ├── 05-httproute.yaml
│   └── 06-hpa.yaml
└── test/
    └── locustfile.py
```

### 7. Código mínimo funcional de la API HTTP

**Archivo:** `api-rust/src/main.rs`**Explicación:** Expone dos rutas: `/` (GET) obligatoria para que el Balanceador de Carga de GCP marque el servicio como "Saludable" (Health Check), y `/grpc-2026` (POST) para recibir el JSON. Incluye un bucle matemático para forzar el HPA intencionalmente.

### 8. Dockerfile

**Archivo:** `api-rust/Dockerfile`
Compilación multi-etapa. Genera una imagen final de Linux Debian súper ligera (~50MB) ideal para la nube.

### 9. Script de Locust

**Archivo:** `test/locustfile.py`
Envía peticiones POST infinitas simulando reportes de guerra. El `host` se pasará por consola al momento de ejecutarlo.

### 10. Manifiestos YAML

*(Están todos en la sección final "para copiar y pegar:)")*.

- **namespace:** Crea `mumnk8s`.
- **deployment:** Llama a la imagen en Docker Hub. Solicita `10m` de CPU.
- **service:** Expone el pod. **Importante:** en GKE, los services detrás de un Gateway deben ser tipo `ClusterIP` nativo o `NodePort`, usaremos `ClusterIP`.
- **gateway:** Usa la clase `gke-l7-global-external-managed` (crea un Load Balancer HTTP externo en GCP).
- **httproute:** Enruta el path `/grpc-2026.`. API: v1
- **hpa:** Escala hasta 3 réplicas al pasar el 30% de CPU.

### 11. Comandos exactos para todo el proceso

```bash
# 1. Crear el clúster en GKE con Gateway API habilitado
gcloud container clusters create mumnk8s-cluster \
    --zone us-central1-a \
    --num-nodes 3 \
    --machine-type e2-standard-2 \
    --gateway-api standard \
    --disk-size=50

# 2. Conectar kubectl al nuevo clúster
gcloud container clusters get-credentials mumnk8s-cluster --zone us-central1-a

gcloud services enable container.googleapis.com compute.googleapis.com logging.googleapis.com

# 3. Construir y subir imagen (Reemplaza TU_USUARIO_DOCKER con tu usuario real)
cd demo1-gcp/api-rust
docker build -t joselorenzana272/api-rust:v1 .
docker push joselorenzana272/api-rust:v1

# 4. Modificar el deployment.yaml para poner tu usuario, luego aplicar YAMLs
cd ../k8s
kubectl apply -f .

# 5. Obtener la IP Pública
kubectl get gateway api-gateway -n mumnk8s -w
```

### 12. Orden exacto de ejecución paso a paso desde cero

1. Crear el clúster de GKE con la bandera `-gateway-api=standard`. (NO OLVIDARLA)
2. Escribir el código en Rust y el Dockerfile.
3. Construir la imagen y subirla a Docker Hub. (Puedes usar la de Josesin que está en Dockerhub ya)
4. Aplicar todos los manifiestos de Kubernetes en GKE.
5. Esperar ~3-5 minutos a que Google aprovisione la IP pública del Gateway.
6. Probar la API con cURL.
7. Arrancar Locust apuntando a la IP pública de GCP.
8. Monitorear el escalamiento del HPA.

### 1. Cómo instalar o habilitar metrics-server

En GKE **no tienes que hacer nada**. El `metrics-server` viene instalado, gestionado y habilitado por defecto por Google. El HPA funciona desde el primer segundo.

### 14. Cómo instalar o habilitar Gateway API en el entorno elegido

Se habilita al momento de crear el clúster agregando la bandera `--gateway-api=standard`. Esto instala los CRDs automáticamente y levanta el controlador `gke-l7-global-external-managed` en el background de GCP.

### 15. Cómo probar la API antes de exponerla con Gateway API

Para asegurar de que el pod de GKE funciona y no es un problema de red de Google:

```bash
kubectl port-forward svc/api-rust-service 8080:80 -n mumnk8s
```

En otra terminal: (esto no es con la IP Pública, es solo un test)

```bash
curl -X POST http://localhost:80/grpc-2026 -H "Content-Type: application/json" -d '{"country": "USA", "warplanes_in_air": 10, "warships_in_water": 5, "timestamp": "2026-03-12T20:15:30Z"}'
```

### 16. Cómo probar la ruta ya expuesta por Gateway API

Ejecuta:

```bash
kubectl get gateway api-gateway -n mumnk8s
```

Buscar la columna `ADDRESS` (ej. `34.120.x.x`). **Atención:** En GCP, un Balanceador demora entre 3 y 5 minutos en enrutar tráfico correctamente.

```bash
curl -X POST http://34.120.x.x/grpc-2026 -H "Content-Type: application/json" -d '{"country": "USA", "warplanes_in_air": 10, "warships_in_water": 5, "timestamp": "2026-03-12T20:15:30Z"}'
```

### 17. Cómo ejecutar Locust

En una terminal nueva:

```bash
cd demo1-gcp/test
python3 -m venv venv
source venv/bin/activate
pip install locust
locust -f locustfile.py --host=http://34.120.x.x  # REEMPLAZA CON LA IP DEL GATEWAY
```

Abrir `http://localhost:8089`. Poner 100 usuarios, 10 de spawn rate y observa las peticiones viajando a la nube. (Lo que quieran realmente)

### 18. Cómo forzar o facilitar que el HPA sí escale

Usamos un `request` de CPU extremadamente bajo en el pod de Kubernetes (`10m`, 1% de un núcleo). Al mismo tiempo, en el código Rust hay un ciclo `for` matemático temporal. Esta combinación hace que la más mínima petición de Locust rompa la barrera del 30% del HPA instantáneamente.

### 19. Comandos para observar el escalamiento en vivo

Abrir dos terminales lado a lado y ejecuta:

```bash
# Terminal 1: Observar el % de CPU en vivo
kubectl get hpa api-rust-hpa -n mumnk8s -w

# Terminal 2: Observa los pods creándose en GKE
kubectl get pods -n mumnk8s -w

# ver en vivo
kubectl logs -f deployment/api-rust-deploy -n mumnk8s
```

### 20. Qué salida o comportamiento deberían de ver si todo va bien

El Load Balancer responderá 200 OK en los logs de Locust. En GKE, el `TARGET` del HPA subirá de `0%/30%` a algo como `200%/30%`. Luego, verás que `REPLICAS` pasa de 1 a 3.
---

### Archivos:

**`demo1-gcp/api-rust/Cargo.toml`**

```toml
[package]
name = "api-rust"
version = "0.1.0"
edition = "2024"

[dependencies]
axum = "0.7.4"
tokio = { version = "1.36.0", features =["full"] }
serde = { version = "1.0", features = ["derive"] }
serde_json = "1.0"
```

**`demo1-gcp/api-rust/src/main.rs`**

```rust
use axum::{
    routing::{get, post},
    Json, Router,
};
use serde::Deserialize;
use tokio::net::TcpListener;

#[derive(Deserialize, Debug)]
struct Report {
    country: String,
    warplanes_in_air: i32,
    warships_in_water: i32,
    timestamp: String,
}

// Ruta necesaria para el Health Check de Google Cloud Load Balancer
async fn health_check() -> &'static str {
    "OK"
}

// Ruta principal que recibe datos de Locust
async fn handle_report(Json(payload): Json<Report>) -> String {
    // Simulador de carga para que el HPA se active en clase
    let mut _dummy: u64 = 0;
    for i in 0..5_000_000 {
        _dummy = _dummy.wrapping_add(i);
    }

    println!("Recibido reporte de: {}", payload.country);
    format!("Reporte de {} recibido en GKE", payload.country)
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(health_check)) // <-- Crítico para GKE
        .route("/grpc-2026", post(handle_report));

    let listener = TcpListener::bind("0.0.0.0:8080").await.unwrap();
    println!("API Rust en puerto 8080...");
    axum::serve(listener, app).await.unwrap();
}
```

**`demo1-gcp/api-rust/Dockerfile`**

```docker
FROM rust:latest AS builder
WORKDIR /app
COPY Cargo.toml Cargo.lock ./
COPY src ./src
RUN cargo build --release

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/target/release/api-rust /usr/local/bin/api-rust
EXPOSE 8080
CMD ["api-rust"]
```

**`demo1-gcp/test/locustfile.py`**

```python
from locust import HttpUser, task, between
import random

class MilitaryReportUser(HttpUser):
    wait_time = between(0.1, 0.5)

    @task
    def send_report(self):
        countries =["USA", "RUS", "CHN", "ESP", "GMT"]
        payload = {
            "country": random.choice(countries),
            "warplanes_in_air": random.randint(0, 50),
            "warships_in_water": random.randint(0, 30),
            "timestamp": "2026-03-12T20:15:30Z"
        }
        self.client.post("/grpc-2026", json=payload)
```

**`demo1-gcp/k8s/01-namespace.yaml`**

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: mumnk8s
```

**`demo1-gcp/k8s/02-deployment.yaml`**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-rust-deploy
  namespace: mumnk8s
spec:
  replicas: 1
  selector:
    matchLabels:
      app: api-rust
  template:
    metadata:
      labels:
        app: api-rust
    spec:
      containers:
        - name: api-rust
          image: joselorenzana272/api-rust:v1
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
          resources:
            requests:
              cpu: "10m"
              memory: "32Mi"
            limits:
              cpu: "50m"
              memory: "64Mi"
```

**`demo1-gcp/k8s/03-service.yaml`**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api-rust-service
  namespace: mumnk8s
spec:
  selector:
    app: api-rust
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
```

**`demo1-gcp/k8s/04-gateway.yaml`**

```yaml
apiVersion: gateway.networking.k8s.io/v1beta1
kind: Gateway
metadata:
  name: api-gateway
  namespace: mumnk8s
spec:
  gatewayClassName: gke-l7-global-external-managed
  listeners:
  - name: http
    protocol: HTTP
    port: 80
```

**`demo1-gcp/k8s/05-httproute.yaml`**

```yaml
apiVersion: gateway.networking.k8s.io/v1beta1
kind: HTTPRoute
metadata:
  name: api-route
  namespace: mumnk8s
spec:
  parentRefs:
  - name: api-gateway
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /
    backendRefs:
    - name: api-rust-service
      port: 80
```

*(Nota sobre HTTPRoute: En GKE mandamos el prefijo `/` para que deje pasar tanto el Health Check `/` como `/grpc-2026` hacia el mismo servicio de Rust).*

**`demo1-gcp/k8s/06-hpa.yaml`**

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: api-rust-hpa
  namespace: mumnk8s
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: api-rust-deploy
  minReplicas: 1
  maxReplicas: 3
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 30
```

---

🎹
