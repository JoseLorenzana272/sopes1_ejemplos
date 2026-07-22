
# ¿Qué es containerd?

**containerd** es un runtime de contenedores de nivel industrial que maneja el ciclo de vida completo de los contenedores en un sistema host. Es el componente principal que Docker utiliza internamente, pero también puede usarse de forma independiente.

Características principales de containerd:
- **Runtime estándar**: Implementa las especificaciones OCI Runtime y Image
- **Gestión de imágenes**: Descarga, almacena y gestiona imágenes de contenedores
- **Ciclo de vida**: Crea, ejecuta, detiene y elimina contenedores
- **Snapshots**: Maneja sistemas de archivos en capas para los contenedores
- **Networking**: Proporciona capacidades básicas de red para contenedores

**ctr** es la herramienta de línea de comandos que viene con containerd, similar a como `docker` es el cliente para Docker Engine.


---

# Comandos basicos de ctr (CONTAINERD)
### 🔧 1. **Verificar containerd está corriendo (opcional si tienes dudas)**

> **¿Aún no tienes instalado containerd?**
> Consulta el script de instalación aquí: [instalar-containerd.sh](../recursos-adicionales/instalar-containerd.sh)


```bash
sudo systemctl status containerd
```

```bash
ctr --version
```


---

### 📦 2. **Listar imágenes disponibles**

```bash
sudo ctr images ls

# si el comando anterior no funciona prueba con 
sudo ctr images list
```

---

### 📥 3. **Descargar una imagen desde un registry (como Docker Hub)**

```bash
sudo ctr images pull docker.io/library/hello-world:latest
```

**Desde un registro privado como Zot en una maquina virtual**
```bash
sudo ctr images pull --plain-http <IP_VM1_DOCKER>:5000/fiber-api-go:v1
```
---

### 🚀 4. **Ejecutar un contenedor**

```bash
sudo ctr run -t --rm docker.io/library/hello-world:latest my-hello
```

* `-t`: modo interactivo
* `--rm`: elimina el contenedor al salir
* `my-hello`: ID local para el contenedor, no puede conteneder dos puntos



### 🕹️ 4.1. **Ejecutar el contenedor en segundo plano con red de host**

Para ejecutar un contenedor en segundo plano (modo *detached*) y que utilice la red del host, usa el siguiente comando. Debes reemplazar `<IP_VM_DOCKER>` por la IP correspondiente de tu máquina virtual o servidor donde está el registro de imágenes:

```bash
sudo ctr run -d --net-host <IP_VM_DOCKER>:5000/fiber-api-go:v1 my-fiber-api-go
```

- `-d`: ejecuta el contenedor en segundo plano (detached).
- `--tty`: asigna una terminal al contenedor.
- `--plain-http`: utiliza la red del host, permitiendo que el contenedor acceda directamente a los puertos y servicios del host.
- `<IP_VM_DOCKER>:5000/fiber-api-go:v1`: imagen a ejecutar, obtenida desde el registro privado.
- `my-fiber-api-go`: nombre local para el contenedor.

Esto es útil cuando necesitas que el contenedor tenga acceso completo a la red del host, por ejemplo, para exponer servicios en los mismos puertos que el sistema anfitrión.

---

### 📋 5. **Listar contenedores activos**

```bash
sudo ctr containers ls
```

---

### 🔄 6. **Crear y ejecutar un contenedor en pasos separados**

```bash
sudo ctr containers create docker.io/library/alpine:latest my-alpine
sudo ctr tasks start -t my-alpine
```

---

### 📤 7. **Subir una imagen a un registry (como Zot o Docker Hub)**

```bash
sudo ctr images tag docker.io/library/hello-world:latest localhost:5000/hello-world:latest
sudo ctr images push <IP_VM1_DOCKER>:5000/hello-world:latest
```

---

### 🧹 8. **Eliminar una imagen**

```bash
sudo ctr images rm docker.io/library/hello-world:latest
```

---

### 🔧 9. **Eliminar un contenedor**

```bash
sudo ctr containers delete my-hello
```

---

### 💡 10. **Acceder a un shell en un contenedor**

Si el contenedor está corriendo y tiene `/bin/sh`:

```bash
sudo ctr tasks exec -t --exec-id myexecid my-alpine /bin/sh
```

---


### 💀 11. **Eliminar todo**

```bash
sudo ctr tasks list -q | xargs -I {} sudo ctr task kill {}
sudo ctr containers list -q | xargs -I {} sudo ctr containers delete {}
sudo ctr images list -q | xargs -I {} sudo ctr images remove {}
```
