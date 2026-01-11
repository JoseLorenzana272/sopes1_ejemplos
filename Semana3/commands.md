# Semana 3: Dominando la CLI en Docker

> **Objetivo:** Perder el miedo a la pantalla negra. Aprenderemos a navegar, instalar software, gestionar procesos y permisos dentro de un entorno Linux real aislado (un contenedor).

---

## Paso 1: El Entorno de Pruebas (Docker)
Primero, necesitamos un sistema Linux donde podamos romper cosas sin miedo. Usaremos una imagen de **Ubuntu** para tener un sistema operativo completo.

```bash
docker run --rm -it --name mi-servidor-linux ubuntu:latest bash
```


---

## Paso 2: Reconocimiento (Navegación y Archivos)


Estás en un lugar desconocido. ¿Dónde estamos? ¿Qué hay aquí?

```bash
# 1. ¿Dónde estoy? (Print Working Directory)
pwd

# 2. Listar archivos (El clásico)
ls

# 3. Listar archivos ocultos y detalles (El "no tan conocido" pero vital)
# -l: formato de lista (permisos, tamaños)
# -a: all (muestra archivos ocultos que empiezan con punto)
ls -la

# 4. Navegar a la carpeta de configuraciones
cd /etc
ls | head -n 5  # "head" es un truco para ver solo las primeras 5 líneas y no llenar la pantalla

```

---

## Paso 3: Gestión de Paquetes (Instalando Herramientas)


Las imágenes de Docker vienen "peladas" (vacías) por seguridad y ligereza. Vamos a intentar usar un editor de texto.

```bash
# Intenta abrir nano (probablemente fallará)
nano

# ¡Error! "command not found". Como SysAdmin, debes instalarlo.
# Actualizamos los repositorios (SIEMPRE haz esto primero en apt)
apt update

# Instalamos un editor (nano) y un monitor de procesos (htop)
# -y: Acepta automáticamente para no preguntar "¿Desea continuar?"
apt install -y nano htop procps

```

---

## Paso 4: Usuarios y Permisos (El "coco" de Linux)


Vamos a crear un script secreto y protegerlo.

1. **Crear el archivo:**
```bash
cd /home
# 'echo' mete texto dentro de un archivo nuevo
echo "echo '¡Hola Hackers de SO1!'" > secreto.sh

```


2. **Verificar permisos:**
```bash
ls -l secreto.sh
# Verás algo como "-rw-r--r--". Faltan las "x" (eXecute).

```


3. **Intentar ejecutarlo (y fallar):**
```bash
./secreto.sh
# Resultado: Permission denied

```


4. **Arreglarlo (chmod):**
```bash
# Damos permisos de ejecución (+x) al dueño
chmod +x secreto.sh

# Verificamos de nuevo (ahora se verá verde o con 'x')
ls -l secreto.sh

# ¡Ejecutamos!
./secreto.sh

```



---

## Paso 5: Procesos y Monitoreo


¿Qué está consumiendo mi memoria?

```bash
# 1. Ver procesos corriendo actualmente (Instantánea)
# a: all users, u: user format, x: procesos sin terminal
ps aux

# 2. Ver procesos en tiempo real (Interactivo - ¡Muy útil!)
htop
# (Para salir de htop, presiona F10 o Ctrl+C)

```

---

## Paso 6: Limpieza y Salida

```bash
# Salir del contenedor (al tener el flag --rm al inicio, el contenedor se autodestruye)
exit

```

---

### 💡 Resumen de Comandos Clave de Hoy

| Comando | Descripción | ¿Por qué importa en Docker? |
| --- | --- | --- |
| `ls -la` | Lista todo (+ocultos) | Para ver archivos `.env` o configuraciones ocultas. |
| `apt update && apt install` | Instala soft | Imprescindible al crear `Dockerfiles`. |
| `chmod +x` | Da ejecución | Necesario para `entrypoint.sh` scripts de inicio. |
| `ps aux` | Muestra procesos | Para debuggear por qué tu servidor se trabó. |
| `cat archivo` | Muestra contenido | Para leer logs rápidamente. |



***
