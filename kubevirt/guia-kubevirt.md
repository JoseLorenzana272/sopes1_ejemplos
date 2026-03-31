# KubeVirt + Grafana en Alpine - Ejemplo extra

## Requisitos previos
- Cluster GKE creado con suficientes recursos
- KubeVirt instalado y todos los pods en `Running`
- `virtctl` instalado***

---

## Paso 1: Crear el Cluster GKE

```bash
gcloud container clusters create mumnk8s-cluster \
    --zone us-central1-a \
    --num-nodes 3 \
    --machine-type e2-standard-4 \
    --gateway-api standard \
    --disk-size=100
```

> Usar `e2-standard-4` y `--disk-size=100` es importante. Con `e2-standard-2` o 50GB no hay suficientes recursos.

---

## Paso 2: Instalar KubeVirt

```bash
# Instalar el operador
kubectl apply -f https://github.com/kubevirt/kubevirt/releases/download/v1.7.2/kubevirt-operator.yaml

# Esperar que el operador esté listo
kubectl rollout status deployment/virt-operator -n kubevirt --timeout=300s

# Parchear para GKE (quita restricciones de nodos maestros)
kubectl patch deployment virt-operator -n kubevirt --type="json" \
  -p='[{"op": "remove", "path": "/spec/template/spec/affinity"}, {"op": "remove", "path": "/spec/template/spec/tolerations"}]'

# Instalar KubeVirt con emulación por software (necesario en GKE)
**emul-kubevirt.yaml** (kubectl apply -f)
apiVersion: kubevirt.io/v1
kind: KubeVirt
metadata:
  name: kubevirt
  namespace: kubevirt
spec:
  certificateRotateStrategy: {}
  configuration:
    developerConfiguration:
      useEmulation: true
  customizeComponents: {}
  imagePullPolicy: IfNotPresent
  workloadUpdateStrategy: {}
  infra:
    nodePlacement:
      tolerations:
      - key: "CriticalAddonsOnly"
        operator: "Exists"
  workloads:
    nodePlacement: {}

# Esperar que todos los pods estén Running
kubectl get pods -n kubevirt -w
```

---

## Paso 3: Instalar virtctl 
(si ya lo tienen instalado les dará error, para arreglarlo solo deben borrar virtctl y volver a instalarlo)
```bash
VERSION=$(kubectl get kubevirt.kubevirt.io/kubevirt -n kubevirt -o=jsonpath="{.status.observedKubeVirtVersion}")
ARCH=$(uname -s | tr A-Z a-z)-$(uname -m | sed 's/x86_64/amd64/')
curl -L -o virtctl https://github.com/kubevirt/kubevirt/releases/download/${VERSION}/virtctl-${VERSION}-${ARCH}
chmod +x virtctl
sudo mv virtctl /usr/local/bin/
```

---

## Paso 4: Crear la VM Alpine con disco extra
alpine-vm.yaml (kubectl apply -f)
```yaml

apiVersion: kubevirt.io/v1
kind: VirtualMachine
metadata:
  name: alpine-test
spec:
  runStrategy: Always
  template:
    spec:
      domain:
        devices:
          disks:
            - name: containerdisk
              disk:
                bus: virtio
            - name: emptydisk
              disk:
                bus: virtio
          interfaces:
            - name: default
              masquerade: {}
        resources:
          requests:
            memory: 2Gi
      networks:
        - name: default
          pod: {}
      volumes:
        - name: containerdisk
          containerDisk:
            image: quay.io/kubevirt/alpine-container-disk-demo:devel
        - name: emptydisk
          emptyDisk:
            capacity: 5Gi

```

Monitorear hasta que esté `Running`:
```bash
kubectl get vmi alpine-test -w
```

---

## Paso 5: Entrar a la VM e instalar Grafana

```bash
virtctl console alpine-test
# usuario: root (sin contraseña)
```

Dentro de la VM como root:

```sh
# 1. Levantar red (necesario cada vez que reinicia la VM)
ip link set eth0 up && udhcpc -i eth0

# 2. Formatear el disco extra de 5GB (solo la primera vez)
mkdosfs -F 32 /dev/vdb
mkdir /mnt/data
mount /dev/vdb /mnt/data

# 3. Descargar Grafana al disco extra
wget https://dl.grafana.com/oss/release/grafana-11.5.0.linux-amd64.tar.gz -O /mnt/data/grafana.tar.gz

# 4. Extraer
tar -zxvf /mnt/data/grafana.tar.gz -C /mnt/data/

# 5. Crear directorios de datos
mkdir -p /mnt/data/grafana-data /mnt/data/grafana-logs /mnt/data/grafana-plugins

# 6. Arrancar Grafana
/mnt/data/grafana-v11.5.0/bin/grafana-server \
  --homepath /mnt/data/grafana-v11.5.0 \
  cfg:default.paths.data=/mnt/data/grafana-data \
  cfg:default.paths.logs=/mnt/data/grafana-logs \
  cfg:default.paths.plugins=/mnt/data/grafana-plugins &

# 7. Verificar que corre en el puerto 3000
netstat -tlnp | grep 3000
```

Salir de la consola: `Ctrl + ]`

> **Nota:** El disco `/dev/vdb` (FAT32) y la red se pierden al reiniciar la VM. Ejecutar los pasos 1 y 2 del mount + red cada vez que reinicies, y el paso 6 para volver a arrancar Grafana.

---

## Paso 6: Exponer Grafana con NodePort
service-grafana.yaml (kubectl apply -f)
```yaml
apiVersion: v1
kind: Service
metadata:
  name: grafana-nodeport
spec:
  type: NodePort
  selector:
    vm.kubevirt.io/name: alpine-test
  ports:
    - name: grafana
      protocol: TCP
      port: 3000
      targetPort: 3000
      nodePort: 32000
```

Verificar que tiene Endpoints:
```bash
kubectl describe svc grafana-nodeport | grep Endpoints
# Debe mostrar: Endpoints: <IP>:3000
```

---

## Paso 7: Abrir el firewall en GCP

```bash
gcloud compute firewall-rules create grafana-nodeport \
  --allow tcp:32000 \
  --source-ranges 0.0.0.0/0 \
  --description "Grafana NodePort"
```

---

## Paso 8: Acceder a Grafana

Obtén la IP externa de cualquier nodo:
```bash
kubectl get nodes -o wide
# Columna EXTERNAL-IP
```

Abre en el navegador:
```
http://<EXTERNAL-IP>:32000
```

Credenciales por defecto: **admin / admin**

---

## Resumen de IPs y puertos

| Componente | Puerto interno | Puerto externo |
|------------|---------------|----------------|
| Grafana VM | 3000          | 32000 (NodePort) |

