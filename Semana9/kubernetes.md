### **Ejemplo 2: Kubernetes en Minikube (Despliegue Real)**

Ya que tienes Minikube, haremos algo visual: **Desplegar una app y escalarla.**

#### **Paso 1: Iniciar Minikube**

En tu terminal:

```bash
minikube start

```

#### **Paso 2: Manifiesto Declarativo (YAML)**

Crea un archivo en VS Code llamado `k8s-demo.yaml`.
Vamos a desplegar una imagen de "echo-server" (que devuelve info del pod) para que se note cuando cambiamos de réplica.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echo-deployment
spec:
  replicas: 2  # <--- Empezamos con 2
  selector:
    matchLabels:
      app: echo
  template:
    metadata:
      labels:
        app: echo
    spec:
      containers:
      - name: echo
        image: k8s.gcr.io/echoserver:1.4
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: echo-service
spec:
  selector:
    app: echo
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: LoadBalancer

```

#### **Paso 3: Aplicar y Explicar**

Ejecuta:

```bash
kubectl apply -f k8s-demo.yaml

```

* **Verificar:** `kubectl get pods` (Verás 2 pods naciendo).
* **Explicar:** "Yo no creé los pods. Yo le dije a K8s 'quiero 2 réplicas' (Deployment) y él las creó por mí."

#### **Paso 4: Verlo en el navegador (Tunnel)**

Como Minikube es local, necesitamos un túnel para ver la IP externa:

```bash
minikube tunnel

```

*(Deja esta terminal corriendo, pedirá sudo).*

Ahora, en otra terminal busca la IP externa:

```bash
kubectl get svc echo-service

```

Abre esa IP en tu navegador (o `localhost` si tunnel funcionó bien). Verás datos técnicos del servidor.

#### **Paso 5: El Gran Final (Escalado y Auto-Healing)**

1. **Escalar:**
Cambia el YAML a `replicas: 10` y aplica de nuevo: `kubectl apply -f k8s-demo.yaml`.
* Muestra `kubectl get pods`. ¡Bum! 10 servidores arriba en segundos.


2. **Matar (Chaos Engineering):**
Borra un pod al azar: `kubectl delete pod echo-deployment-xxxx`.
* Muestra `kubectl get pods` inmediatamente.
* **Lección:** *"Kubernetes notó que faltaba uno y creó el reemplazo antes de que parpadearan. Eso es orquestación."*



### Resumen

1. **gRPC:** Programación pura y comunicación entre lenguajes.
2. **K8s:** Infraestructura como código y alta disponibilidad en vivo.