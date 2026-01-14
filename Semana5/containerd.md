# Manual Técnico

## Operando a Corazón Abierto: Containerd vs. Docker

### Objetivo

Demostrar que Docker es una herramienta de alto nivel que abstrae containerd, y que los contenedores se ejecutan dentro de *namespaces*. Asimismo, evidenciar que `ctr` (CLI de containerd) requiere comandos explícitos y no aplica automatismos.

---

## Requisitos Previos

* Máquina virtual Linux (la misma usada en prácticas anteriores).
* Docker y containerd instalados.
* Acceso como usuario root (ctr requiere privilegios elevados).

---

## Preparación del Entorno

1. Acceder a la máquina virtual.
2. Obtener privilegios de superusuario:

```bash
sudo su
```

---

## Paso 1: Verificación de Namespaces (Choque de Realidad)

### 1.1 Ejecutar un contenedor con Docker

```bash
docker run -d --name soy-docker alpine sleep 1000
```

Este comando crea y ejecuta un contenedor usando Docker, que internamente utiliza containerd.

### 1.2 Listar contenedores usando ctr (namespace por defecto)

```bash
ctr containers ls
```

Resultado esperado:

* No se muestra ningún contenedor.

### 1.3 Explicación técnica

* Docker **no utiliza el namespace por defecto de containerd**.
* Docker ejecuta sus contenedores dentro del namespace **moby**.
* `ctr`, por defecto, opera en el namespace **default**.

### 1.4 Listar contenedores en el namespace correcto

```bash
ctr -n moby containers ls
```

Resultado esperado:

* El contenedor `soy-docker` aparece listado.

### Conclusión del paso

Docker actúa como un administrador que organiza los contenedores de containerd dentro de un namespace específico llamado `moby`.

---

## Paso 2: Descarga Manual de Imágenes (Pull Explícito)

### 2.1 Intento de descarga usando un nombre incompleto

```bash
ctr images pull alpine
```

Resultado esperado:

* Error indicando que la referencia de imagen no fue encontrada.

### 2.2 Explicación técnica

* Docker asume automáticamente Docker Hub (`docker.io`) y la librería `library`.
* Containerd **no asume valores por defecto**.
* Se debe usar el nombre completamente calificado (FQDN).

### 2.3 Descarga correcta de la imagen

```bash
ctr images pull docker.io/library/alpine:latest
```

Resultado esperado:

* Descarga exitosa de las capas de la imagen.

---

## Paso 3: Ejecución Manual de un Contenedor

### 3.1 Ejecutar un contenedor con ctr

Estructura general del comando:

```bash
ctr run [flags] [imagen] [id-contenedor] [comando]
```

Ejecutar el contenedor:

```bash
ctr run -d docker.io/library/alpine:latest demo-bajo-nivel sh -c "echo 'Hola desde las entrañas de Containerd' && sleep 1000"
```

Notas técnicas:

* El ID del contenedor debe definirse manualmente.
* No existe autogeneración de nombres.
* Se ejecuta directamente sobre containerd.

### 3.2 Verificar contenedores en containerd

```bash
ctr containers ls
```

Resultado esperado:

* El contenedor `demo-bajo-nivel` aparece listado.

### 3.3 Verificar visibilidad desde Docker

```bash
docker ps
```

Resultado esperado:

* El contenedor **no aparece**.

### Conclusión del paso

El contenedor existe en el sistema operativo y en containerd, pero Docker no lo reconoce porque no fue creado en su namespace.

---

## Paso 4: Interacción con el Contenedor (Exec)

### 4.1 Acceder al contenedor

```bash
ctr tasks exec -t --exec-id shell1 demo-bajo-nivel sh
```

Resultado esperado:

* Acceso a una shell dentro del contenedor (`/ #`).

### 4.2 Verificar procesos internos

Dentro del contenedor, ejecutar:

```bash
ps aux
```

Resultado esperado:

* El proceso `sleep` se encuentra activo.

### 4.3 Salir del contenedor

```bash
exit
```

---

## Paso 5: Limpieza Manual de Recursos

### 5.1 Listar tareas activas

```bash
ctr tasks ls
```

### 5.2 Detener la tarea (proceso)

```bash
ctr tasks kill demo-bajo-nivel
```

### 5.3 Eliminar el contenedor (metadata)

```bash
ctr containers delete demo-bajo-nivel
```

Notas:

* Containerd no elimina recursos automáticamente.
* La limpieza debe realizarse explícitamente.

---

## Resumen Conceptual (Alto Nivel vs Bajo Nivel)

### Resolución de nombres

* Docker: asume `docker.io/library/imagen`.
* Containerd: requiere nombre completo (`docker.io/library/...`).

### Namespaces

* Docker opera en el namespace `moby`.
* Containerd usa `default`.
* Esto permite que Docker y Kubernetes coexistan sin conflictos.

### Nivel de abstracción

* Docker: orientado a usuarios humanos.
* Containerd: orientado a sistemas, automatización y programadores.

### Flujo de ejecución

* Docker: `docker run`
* Containerd: `pull → create → start task`

---

## Conclusión Final

Este procedimiento demuestra claramente la diferencia entre una herramienta de alto nivel (Docker) y el motor real de contenedores (containerd), evidenciando cómo Docker simplifica y oculta complejidad que containerd expone de forma explícita.

---
