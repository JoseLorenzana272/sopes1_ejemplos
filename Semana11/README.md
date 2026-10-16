# Operator Helm-based: ValkeyInstance


```
 kubectl apply ValkeyInstance  ──►  helm-operator (Deployment)
                                          │  lee watches.yaml
                                          ▼
                                helm-charts/valkeyinstance
                                          │  helm install/upgrade/uninstall
                                          ▼
                              StatefulSet + Service + PVC
```

- El **CRD** define el nuevo tipo `ValkeyInstance`.
- `watches.yaml` enlaza ese tipo con un **chart**.
- El `spec` del CR se convierte en los **values** del chart.
- Crear / editar / borrar el CR = `helm install` / `upgrade` / `uninstall` (reconciliación).
- Nombre del release = nombre del CR (por eso el chart usa `.Release.Name`).

## Estructura

```
.
├── Dockerfile
├── watches.yaml
├── helm-charts/valkeyinstance/   # Chart.yaml, values.yaml, templates/
├── deploy/                       # CRD, RBAC, Deployment del operator
└── samples/valkeyinstance.yaml   # el Custom Resource
```

## Cómo generarlo con operator-sdk

```bash
mkdir valkey-operator && cd valkey-operator
operator-sdk init --plugins helm --domain example.com
operator-sdk create api --group cache --version v1alpha1 --kind ValkeyInstance
# Esto genera helm-charts/valkeyinstance con un chart de ejemplo (nginx):
# reemplacen sus templates por los de Valkey.
```

## Construir, publicar en Zot y desplegar

```bash
# 1. Build y push a Zot (HTTPS)
docker build -t <IP-ZOT>:5000/valkey-operator:v0.1.0 .
docker push  <IP-ZOT>:5000/valkey-operator:v0.1.0

# 2. Instalar el operator en GKE
kubectl apply -f deploy/01-crd.yaml
kubectl apply -f deploy/02-rbac.yaml
kubectl apply -f deploy/03-operator.yaml
kubectl -n mumn get pods -l app=valkey-operator

# 3. Crear el Custom Resource
kubectl apply -f samples/valkeyinstance.yaml
kubectl -n mumn get statefulset,svc,pvc
```

## Probar la reconciliación

```bash
# Editar el CR -> el operator hace helm upgrade
kubectl -n mumn patch valkeyinstance valkey --type merge \
  -p '{"spec":{"maxmemory":"256mb"}}'

# Probar Valkey
kubectl -n mumn exec -it valkey-0 -- valkey-cli ping

# Persistencia: borrar el pod y verificar que los datos siguen
kubectl -n mumn exec valkey-0 -- valkey-cli set prueba hola
kubectl -n mumn delete pod valkey-0
kubectl -n mumn exec valkey-0 -- valkey-cli get prueba

# Borrar el CR -> el operator hace helm uninstall
kubectl -n mumn delete valkeyinstance valkey
```

Los consumers se conectan a: `valkey.mumn.svc.cluster.local:6379`
(o `valkey-0.valkey.mumn.svc.cluster.local` por ser headless/StatefulSet).

