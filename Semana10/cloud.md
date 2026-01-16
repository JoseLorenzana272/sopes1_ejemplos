
---

### **⚠️ Advertencia para la clase**

Aunque Google da créditos gratis ($300), recordar: **"Cluster que no se usa, cluster que se borra"**. Un cluster de GKE olvidado puede consumir el saldo rápidamente. Al final de la guía está el comando de destrucción.

---

### **Fase 1: Preparando el Terreno (CLI)**

Asumo que ya tienes el `google-cloud-sdk` instalado. Si no, instálalo rápido.

1. **Loguearse:** (Esto abrirá el navegador)
```bash
gcloud auth login

```


2. **Crear/Seleccionar Proyecto:**
```bash
# Crea un proyecto nuevo para la clase
gcloud projects create sopes1-demo-gke --name="Demo Clase SO1"

# Selecciónalo
gcloud config set project sopes1-demo-gke

```


3. **Habilitar la API de Kubernetes:** (Tarda unos segundos)
```bash
gcloud services enable container.googleapis.com

```



---

### **Fase 2: Naciendo en la Nube (Crear el Cluster)**

Vamos a crear un cluster pequeño para no gastar mucho, pero funcional.

```bash
# Crea un cluster de 1 solo nodo (máquina virtual) en la zona central
# Tarda unos 3-5 minutos en provisionar. Aprovecha para explicar qué es un "Nodo".
gcloud container clusters create cluster-clase \
    --num-nodes=1 \
    --zone=us-central1-a \
    --machine-type=e2-medium

```

* **Diferencia con Minikube:** Aquí Google está buscando un servidor físico en un datacenter de Iowa (`us-central1`) y reservándolo para ti.

Una vez termine, conecta tu `kubectl` local a la nube:

```bash
gcloud container clusters get-credentials cluster-clase --zone us-central1-a

```

---

### **Fase 3: Elasticidad Real (HPA)**

Usaremos la misma imagen de antes (`php-apache`) porque es perfecta para demos de CPU, pero ahora verás cómo GKE maneja el networking.

Crea el archivo `gke-demo.yaml` en VS Code:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: servidor-elastico
spec:
  selector:
    matchLabels:
      run: php-apache
  replicas: 1
  template:
    metadata:
      labels:
        run: php-apache
    spec:
      containers:
      - name: php-apache
        image: registry.k8s.io/hpa-example # Imagen diseñada para estresarse
        ports:
        - containerPort: 80
        resources:
          # IMPORTANTE: El HPA necesita estos límites para calcular porcentajes
          limits:
            cpu: 500m
          requests:
            cpu: 200m
---
apiVersion: v1
kind: Service
metadata:
  name: servidor-elastico
spec:
  ports:
  - port: 80
  selector:
    run: php-apache
  # AQUÍ ESTÁ LA MAGIA DE GKE:
  # Esto provisionará una IP Pública real de Google y un Balanceador de Carga físico.
  type: LoadBalancer

```

1. **Despliégalo:**
```bash
kubectl apply -f gke-demo.yaml

```


2. **Configura el Autoescalado (HPA):**
```bash
# "Si la CPU pasa del 50%, crea réplicas hasta llegar a 10"
kubectl autoscale deployment servidor-elastico --cpu-percent=50 --min=1 --max=10

```



---

### **Fase 4: Ver la IP Pública (Wow Factor)**

En Minikube esto era un dolor de cabeza (túneles). En GKE es nativo.

```bash
kubectl get services

```

* Mira la columna `EXTERNAL-IP`. Al principio dirá `<pending>`.
* Ejecuta el comando varias veces hasta que salga una IP real (ej. `34.12.55.19`).
* **Demo:** Dile a tus alumnos: *"Entren todos a `http://34.12.55.19` desde sus celulares"*. Verán el mensaje "OK!".

---

### **Fase 5: El Ataque DDoS (Stress Test)**

Ahora demostraremos la **Elasticidad** (Tema de Semana 10).

1. **Monitor en vivo (Terminal 1):**
```bash
kubectl get hpa -w

```


2. **Generar Carga (Terminal 2):**
Crearemos un pod "atacante" dentro del cluster que bombardee al servidor.
```bash
kubectl run -i --tty generador-carga --rm --image=busybox:1.28 -- restart=Never -- /bin/sh -c "while true; do wget -q -O- http://servidor-elastico; done"

```



**Lo que verán en clase:**

1. En la Terminal 1, el `TARGETS` subirá a `200%/50%`.
2. GKE dirá: *"Necesito más potencia"*.
3. La columna `REPLICAS` subirá de `1` -> `4` -> `8`.
4. **Explicación:** *"Google Cloud está creando contenedores automáticamente para absorber el tráfico. Si esto fuera una tienda en Black Friday, no se hubiera caído".*

---

### **Fase 6: LA DESTRUCCIÓN (¡Muy Importante!)**

Para evitar cobros sorpresa, termina la clase ejecutando esto frente a ellos (enseña buenas prácticas de costos):

```bash
# 1. Borrar el cluster (Esto borra las VMs, el balanceador y los discos)
gcloud container clusters delete cluster-clase --zone us-central1-a --quiet

# 2. Confirmar que no queda nada
gcloud container clusters list

```
