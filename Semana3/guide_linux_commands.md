# Guía Completa de Comandos Básicos de Linux

## Introducción

Esta guía te ayudará a dominar los comandos esenciales de Linux para navegar y trabajar eficientemente en la terminal. La terminal es una herramienta poderosa que te permite controlar tu sistema operativo de manera directa y eficiente.

---

## 1. Navegación del Sistema de Archivos

### `pwd` - Print Working Directory

Muestra la ruta completa del directorio actual donde te encuentras.

```bash
pwd
```

**Ejemplo de salida:**

```
/home/usuario/documentos
```

**Cuándo usar:** Cuando necesitas saber exactamente dónde estás en el sistema de archivos.

---

### `ls` - List

Lista los archivos y directorios en el directorio actual.

**Sintaxis básica:**

```bash
ls [opciones] [directorio]
```

**Opciones comunes:**

- `ls` - Lista archivos básica
- `ls -l` - Lista detallada (permisos, propietario, tamaño, fecha)
- `ls -a` - Muestra archivos ocultos (los que empiezan con punto)
- `ls -lh` - Lista detallada con tamaños legibles (KB, MB, GB)
- `ls -la` - Lista detallada incluyendo archivos ocultos
- `ls -lt` - Ordena por fecha de modificación (más reciente primero)
- `ls -lS` - Ordena por tamaño (más grande primero)
- `ls -R` - Lista recursiva (incluye subdirectorios)

**Ejemplos:**

```bash
# Listar archivos en el directorio actual
ls

# Listar con detalles
ls -l

# Listar archivos ocultos
ls -a

# Listar con tamaños legibles
ls -lh

# Listar contenido de otro directorio
ls /home/usuario/descargas
```

**Interpretando `ls -l`:**

```
-rw-r--r-- 1 usuario grupo 4096 ene 15 10:30 archivo.txt
│││││││││  │ │       │     │    │            └─ nombre
│││││││││  │ │       │     │    └─ fecha y hora
│││││││││  │ │       │     └─ tamaño en bytes
│││││││││  │ │       └─ grupo propietario
│││││││││  │ └─ usuario propietario
│││││││││  └─ número de enlaces
└─ permisos (tipo y acceso)
```

---

### `cd` - Change Directory

Cambia el directorio actual a otro directorio.

**Sintaxis:**

```bash
cd [directorio]
```

**Atajos especiales:**

- `cd` o `cd ~` - Va al directorio home del usuario
- `cd ..` - Sube un nivel (directorio padre)
- `cd ../..` - Sube dos niveles
- `cd -` - Regresa al directorio anterior
- `cd /` - Va al directorio raíz del sistema
- `cd ./carpeta` - Entra a una carpeta en el directorio actual

**Ejemplos:**

```bash
# Ir al directorio home
cd

# Ir a documentos
cd documentos

# Subir un nivel
cd ..

# Ir usando ruta absoluta
cd /var/log

# Ir usando ruta relativa
cd ../descargas

# Regresar al directorio anterior
cd -
```

**Rutas absolutas vs relativas:**

- **Absoluta:** Empieza desde la raíz `/home/usuario/documentos`
- **Relativa:** Desde tu ubicación actual `../carpeta` o `./archivo`

---

## 2. Manipulación de Archivos y Directorios

### `mkdir` - Make Directory

Crea nuevos directorios (carpetas).

**Sintaxis:**

```bash
mkdir [opciones] nombre_directorio
```

**Opciones útiles:**

- `mkdir -p` - Crea directorios padres si no existen
- `mkdir -v` - Modo verbose (muestra lo que hace)

**Ejemplos:**

```bash
# Crear un directorio simple
mkdir mi_carpeta

# Crear múltiples directorios
mkdir carpeta1 carpeta2 carpeta3

# Crear estructura anidada
mkdir -p proyectos/web/css

# Crear con permisos específicos
mkdir -m 755 publica
```

---

### `touch`

Crea archivos vacíos o actualiza la fecha de modificación de archivos existentes.

**Sintaxis:**

```bash
touch archivo
```

**Ejemplos:**

```bash
# Crear archivo vacío
touch nuevo.txt

# Crear múltiples archivos
touch archivo1.txt archivo2.txt archivo3.txt

# Actualizar fecha de modificación
touch archivo_existente.txt
```

---

### `cp` - Copy

Copia archivos y directorios.

**Sintaxis:**

```bash
cp [opciones] origen destino
```

**Opciones importantes:**

- `cp -r` - Copia directorios recursivamente
- `cp -i` - Modo interactivo (pide confirmación antes de sobrescribir)
- `cp -v` - Modo verbose (muestra lo que copia)
- `cp -u` - Copia solo si el origen es más reciente
- `cp -a` - Modo archivo (preserva permisos, fechas, etc.)

**Ejemplos:**

```bash
# Copiar archivo en el mismo directorio
cp archivo.txt copia_archivo.txt

# Copiar archivo a otro directorio
cp archivo.txt /home/usuario/documentos/

# Copiar directorio completo
cp -r carpeta carpeta_backup

# Copiar con confirmación
cp -i archivo.txt destino/

# Copiar múltiples archivos a un directorio
cp archivo1.txt archivo2.txt directorio_destino/
```

---

### `mv` - Move

Mueve o renombra archivos y directorios.

**Sintaxis:**

```bash
mv [opciones] origen destino
```

**Opciones:**

- `mv -i` - Pide confirmación antes de sobrescribir
- `mv -v` - Muestra lo que mueve
- `mv -n` - No sobrescribe archivos existentes

**Ejemplos:**

```bash
# Renombrar archivo
mv viejo_nombre.txt nuevo_nombre.txt

# Mover archivo a otro directorio
mv archivo.txt /home/usuario/documentos/

# Mover directorio
mv carpeta_antigua carpeta_nueva

# Mover múltiples archivos
mv archivo1.txt archivo2.txt directorio_destino/

# Renombrar directorio
mv directorio_viejo directorio_nuevo
```

---

### `rm` - Remove

Elimina archivos y directorios. **¡PRECAUCIÓN!** Esta operación es permanente.

**Sintaxis:**

```bash
rm [opciones] archivo
```

**Opciones:**

- `rm -r` - Elimina directorios recursivamente
- `rm -f` - Forzar eliminación sin confirmación
- `rm -i` - Pide confirmación para cada archivo
- `rm -rf` - **¡PELIGROSO!** Elimina todo sin preguntar

**Ejemplos:**

```bash
# Eliminar archivo
rm archivo.txt

# Eliminar con confirmación
rm -i documento.pdf

# Eliminar directorio vacío
rmdir carpeta_vacia

# Eliminar directorio con contenido
rm -r carpeta

# Eliminar múltiples archivos
rm archivo1.txt archivo2.txt archivo3.txt

# Eliminar archivos por patrón
rm *.tmp
```

**⚠️ ADVERTENCIA:** Nunca ejecutes `rm -rf /` o `rm -rf /*` - ¡eliminará todo tu sistema!

---

## 3. Visualización de Contenido

### `cat` - Concatenate

Muestra el contenido completo de archivos.

**Sintaxis:**

```bash
cat [opciones] archivo
```

**Opciones:**

- `cat -n` - Muestra números de línea
- `cat -b` - Numera solo líneas no vacías
- `cat -s` - Suprime líneas vacías repetidas

**Ejemplos:**

```bash
# Ver contenido de archivo
cat archivo.txt

# Ver múltiples archivos
cat archivo1.txt archivo2.txt

# Ver con números de línea
cat -n script.py

# Concatenar archivos en uno nuevo
cat archivo1.txt archivo2.txt > combinado.txt
```

---

### `less`

Visualiza archivos página por página (mejor para archivos grandes).

**Sintaxis:**

```bash
less archivo
```

**Controles dentro de less:**

- `Espacio` - Avanza una página
- `b` - Retrocede una página
- `/texto` - Busca "texto" hacia adelante
- `?texto` - Busca "texto" hacia atrás
- `n` - Siguiente resultado de búsqueda
- `N` - Resultado anterior de búsqueda
- `g` - Ir al inicio del archivo
- `G` - Ir al final del archivo
- `q` - Salir

**Ejemplos:**

```bash
# Ver archivo largo
less archivo_grande.log

# Ver con números de línea
less -N archivo.txt
```

---

### `head`

Muestra las primeras líneas de un archivo (por defecto 10).

**Sintaxis:**

```bash
head [opciones] archivo
```

**Opciones:**

- `head -n 20` - Muestra las primeras 20 líneas
- `head -n -5` - Muestra todo excepto las últimas 5 líneas

**Ejemplos:**

```bash
# Ver primeras 10 líneas
head archivo.txt

# Ver primeras 20 líneas
head -n 20 log.txt

# Ver primeras líneas de múltiples archivos
head archivo1.txt archivo2.txt
```

---

### `tail`

Muestra las últimas líneas de un archivo (por defecto 10).

**Sintaxis:**

```bash
tail [opciones] archivo
```

**Opciones:**

- `tail -n 20` - Muestra las últimas 20 líneas
- `tail -f` - Sigue el archivo en tiempo real (útil para logs)
- `tail -F` - Sigue el archivo incluso si se recrea

**Ejemplos:**

```bash
# Ver últimas 10 líneas
tail archivo.txt

# Ver últimas 50 líneas
tail -n 50 log.txt

# Seguir un archivo en tiempo real (Ctrl+C para salir)
tail -f /var/log/syslog
```

---

## 4. Búsqueda y Filtrado

### `find`

Busca archivos y directorios en el sistema de archivos.

**Sintaxis:**

```bash
find [ruta] [opciones] [criterios]
```

**Opciones comunes:**

- `-name` - Busca por nombre
- `-iname` - Busca por nombre (ignorando mayúsculas)
- `-type f` - Solo archivos
- `-type d` - Solo directorios
- `-size` - Por tamaño
- `-mtime` - Por fecha de modificación

**Ejemplos:**

```bash
# Buscar archivo por nombre en directorio actual
find . -name "archivo.txt"

# Buscar todos los archivos .pdf
find /home -name "*.pdf"

# Buscar ignorando mayúsculas
find . -iname "*.JPG"

# Buscar solo directorios
find . -type d -name "proyecto*"

# Buscar archivos mayores a 100MB
find . -type f -size +100M

# Buscar archivos modificados en los últimos 7 días
find . -type f -mtime -7

# Buscar y eliminar archivos .tmp
find . -name "*.tmp" -delete
```

---

### `grep` - Global Regular Expression Print

Busca patrones de texto dentro de archivos.

**Sintaxis:**

```bash
grep [opciones] "patrón" archivo
```

**Opciones útiles:**

- `grep -i` - Ignora mayúsculas/minúsculas
- `grep -r` - Búsqueda recursiva en directorios
- `grep -n` - Muestra números de línea
- `grep -v` - Invierte la búsqueda (líneas que NO coinciden)
- `grep -c` - Cuenta coincidencias
- `grep -w` - Busca palabra completa
- `grep -A 3` - Muestra 3 líneas después de la coincidencia
- `grep -B 3` - Muestra 3 líneas antes de la coincidencia

**Ejemplos:**

```bash
# Buscar palabra en archivo
grep "error" log.txt

# Buscar ignorando mayúsculas
grep -i "warning" log.txt

# Buscar recursivamente en directorio
grep -r "TODO" /home/usuario/proyecto/

# Buscar con números de línea
grep -n "función" script.py

# Contar coincidencias
grep -c "error" log.txt

# Buscar múltiples patrones
grep -e "error" -e "warning" log.txt

# Buscar en múltiples archivos
grep "import" *.py
```

---

### `which`

Muestra la ruta completa de un comando ejecutable.

**Sintaxis:**

```bash
which comando
```

**Ejemplos:**

```bash
# Encontrar ubicación de python
which python

# Encontrar ubicación de múltiples comandos
which python java node
```

---

## 5. Gestión de Permisos

### `chmod` - Change Mode

Cambia los permisos de archivos y directorios.

**Sistema de permisos:**

- `r` (read) = 4 - Leer
- `w` (write) = 2 - Escribir
- `x` (execute) = 1 - Ejecutar

**Grupos de permisos:**

- Usuario propietario (u)
- Grupo (g)
- Otros (o)

**Sintaxis numérica:**

```bash
chmod [permisos] archivo
```

**Permisos comunes:**

- `755` - rwxr-xr-x (ejecutables, scripts)
- `644` - rw-r--r-- (archivos normales)
- `600` - rw------- (archivos privados)
- `777` - rwxrwxrwx (todos los permisos - raramente recomendado)

**Sintaxis simbólica:**

```bash
chmod [u/g/o/a][+/-/=][r/w/x] archivo
```

**Ejemplos:**

```bash
# Dar permisos de ejecución al propietario
chmod u+x script.sh

# Hacer archivo ejecutable para todos
chmod +x programa

# Permisos 644 (lectura/escritura para propietario, solo lectura para otros)
chmod 644 archivo.txt

# Permisos 755 (ejecución para script)
chmod 755 script.sh

# Quitar permisos de escritura para grupo y otros
chmod go-w archivo.txt

# Cambiar permisos recursivamente
chmod -R 755 directorio/
```

---

### `chown` - Change Owner

Cambia el propietario de archivos y directorios.

**Sintaxis:**

```bash
chown [usuario]:[grupo] archivo
```

**Ejemplos:**

```bash
# Cambiar propietario
sudo chown usuario archivo.txt

# Cambiar propietario y grupo
sudo chown usuario:grupo archivo.txt

# Cambiar recursivamente
sudo chown -R usuario:grupo directorio/
```

---

## 6. Información del Sistema

### `whoami`

Muestra tu nombre de usuario actual.

```bash
whoami
```

---

### `hostname`

Muestra el nombre del sistema.

```bash
hostname
```

---

### `uname`

Muestra información del sistema operativo.

**Opciones:**

- `uname -a` - Toda la información
- `uname -r` - Versión del kernel
- `uname -m` - Arquitectura del hardware

**Ejemplos:**

```bash
# Información básica
uname

# Toda la información
uname -a

# Solo versión del kernel
uname -r
```

---

### `df` - Disk Free

Muestra el espacio en disco disponible.

**Sintaxis:**

```bash
df [opciones]
```

**Opciones:**

- `df -h` - Formato legible (GB, MB)
- `df -T` - Muestra tipo de sistema de archivos

**Ejemplos:**

```bash
# Espacio en disco legible
df -h

# Con tipo de sistema de archivos
df -Th
```

---

### `du` - Disk Usage

Muestra el espacio utilizado por archivos y directorios.

**Opciones:**

- `du -h` - Formato legible
- `du -s` - Resumen (solo total)
- `du -a` - Todos los archivos
- `du -d 1` - Profundidad de 1 nivel

**Ejemplos:**

```bash
# Uso del directorio actual
du -h

# Resumen del directorio
du -sh directorio/

# Los 10 directorios más grandes
du -h | sort -rh | head -10
```

---

### `free`

Muestra el uso de memoria RAM.

**Opciones:**

- `free -h` - Formato legible
- `free -m` - En megabytes

**Ejemplos:**

```bash
# Memoria en formato legible
free -h
```

---

### `top`

Muestra procesos en ejecución en tiempo real.

**Controles dentro de top:**

- `q` - Salir
- `k` - Matar proceso (pide PID)
- `M` - Ordenar por uso de memoria
- `P` - Ordenar por uso de CPU
- `h` - Ayuda

**Ejemplo:**

```bash
top
```

**Alternativa moderna:** `htop` (más visual, si está instalado)

---

### `ps` - Process Status

Muestra procesos en ejecución.

**Opciones comunes:**

- `ps aux` - Todos los procesos con detalles
- `ps -ef` - Formato completo

**Ejemplos:**

```bash
# Todos los procesos
ps aux

# Buscar proceso específico
ps aux | grep firefox

# Procesos del usuario actual
ps -u $USER
```

---

## 7. Gestión de Procesos

### `kill`

Termina procesos por su PID (Process ID).

**Sintaxis:**

```bash
kill [señal] PID
```

**Señales comunes:**

- `kill PID` - Terminar normalmente (TERM)
- `kill -9 PID` - Forzar terminación (KILL)
- `kill -15 PID` - Terminar gracefully (TERM)

**Ejemplos:**

```bash
# Terminar proceso con PID 1234
kill 1234

# Forzar terminación
kill -9 1234

# Terminar proceso por nombre
killall firefox
```

---

### `bg` y `fg`

Maneja trabajos en segundo plano y primer plano.

**Ejemplos:**

```bash
# Ejecutar comando en segundo plano
comando &

# Ver trabajos
jobs

# Traer trabajo al primer plano
fg %1

# Enviar trabajo al segundo plano
bg %1

# Suspender proceso actual (Ctrl+Z)
# Luego continuar en segundo plano: bg
```

---

## 8. Compresión y Archivado

### `tar` - Tape Archive

Crea y extrae archivos tar (archivos comprimidos).

**Sintaxis:**

```bash
tar [opciones] archivo.tar archivos
```

**Opciones principales:**

- `c` - Crear archivo
- `x` - Extraer archivo
- `v` - Verbose (muestra progreso)
- `f` - Especifica nombre de archivo
- `z` - Comprimir con gzip (.tar.gz)
- `j` - Comprimir con bzip2 (.tar.bz2)

**Ejemplos:**

```bash
# Crear archivo tar
tar -cvf archivo.tar carpeta/

# Crear archivo tar.gz (comprimido)
tar -czvf archivo.tar.gz carpeta/

# Extraer tar
tar -xvf archivo.tar

# Extraer tar.gz
tar -xzvf archivo.tar.gz

# Ver contenido sin extraer
tar -tvf archivo.tar

# Extraer en directorio específico
tar -xvf archivo.tar -C /ruta/destino/
```

---

### `zip` y `unzip`

Comprime y descomprime archivos ZIP.

**Ejemplos:**

```bash
# Crear archivo zip
zip archivo.zip archivo1.txt archivo2.txt

# Comprimir directorio
zip -r archivo.zip carpeta/

# Extraer zip
unzip archivo.zip

# Extraer en directorio específico
unzip archivo.zip -d /ruta/destino/

# Ver contenido sin extraer
unzip -l archivo.zip
```

---

### `gzip` y `gunzip`

Comprime archivos individuales.

**Ejemplos:**

```bash
# Comprimir archivo (reemplaza original)
gzip archivo.txt

# Descomprimir
gunzip archivo.txt.gz

# Comprimir manteniendo original
gzip -k archivo.txt
```

---

## 9. Redirección y Tuberías

### Redirección de Salida

**Operadores:**

- `>` - Redirige salida (sobrescribe)
- `>>` - Redirige salida (añade al final)
- `2>` - Redirige errores
- `&>` - Redirige salida y errores

**Ejemplos:**

```bash
# Guardar salida en archivo
ls -l > listado.txt

# Añadir al final del archivo
echo "Nueva línea" >> archivo.txt

# Redirigir errores
comando 2> errores.log

# Redirigir todo
comando &> salida_completa.log

# Descartar salida
comando > /dev/null
```

---

### Tuberías (Pipes)

El operador `|` conecta la salida de un comando con la entrada de otro.

**Ejemplos:**

```bash
# Contar líneas en salida
ls -l | wc -l

# Buscar en salida
ps aux | grep firefox

# Ordenar y mostrar primeros resultados
du -h | sort -rh | head -10

# Múltiples tuberías
cat archivo.txt | grep "error" | wc -l

# Buscar y paginar
ls -la | less
```

---

## 10. Utilidades de Texto

### `echo`

Imprime texto en pantalla.

**Ejemplos:**

```bash
# Imprimir texto simple
echo "Hola Mundo"

# Imprimir variable
echo $HOME

# Crear archivo con contenido
echo "Contenido" > archivo.txt

# Añadir línea a archivo
echo "Nueva línea" >> archivo.txt
```

---

### `wc` - Word Count

Cuenta líneas, palabras y caracteres.

**Opciones:**

- `wc -l` - Solo líneas
- `wc -w` - Solo palabras
- `wc -c` - Solo bytes
- `wc -m` - Solo caracteres

**Ejemplos:**

```bash
# Contar todo
wc archivo.txt

# Contar solo líneas
wc -l archivo.txt

# Contar archivos en directorio
ls | wc -l
```

---

### `sort`

Ordena líneas de texto.

**Opciones:**

- `sort -r` - Orden reverso
- `sort -n` - Orden numérico
- `sort -u` - Elimina duplicados

**Ejemplos:**

```bash
# Ordenar archivo
sort archivo.txt

# Ordenar numéricamente
sort -n numeros.txt

# Ordenar y eliminar duplicados
sort -u lista.txt
```

---

### `uniq`

Elimina líneas duplicadas consecutivas.

**Opciones:**

- `uniq -c` - Cuenta repeticiones
- `uniq -d` - Solo muestra duplicados
- `uniq -u` - Solo muestra únicos

**Ejemplos:**

```bash
# Eliminar duplicados (requiere estar ordenado primero)
sort archivo.txt | uniq

# Contar ocurrencias
sort archivo.txt | uniq -c
```

---

## 11. Descarga y Red

### `wget`

Descarga archivos de Internet.

**Sintaxis:**

```bash
wget [opciones] URL
```

**Opciones:**

- `wget -O nombre` - Guardar con nombre específico
- `wget -c` - Continuar descarga interrumpida
- `wget -b` - Descarga en segundo plano

**Ejemplos:**

```bash
# Descargar archivo
wget https://ejemplo.com/archivo.zip

# Descargar con nombre específico
wget -O mi_archivo.zip https://ejemplo.com/archivo.zip

# Continuar descarga
wget -c https://ejemplo.com/archivo_grande.iso
```

---

### `curl`

Transfiere datos desde o hacia un servidor.

**Ejemplos:**

```bash
# Ver contenido de URL
curl https://ejemplo.com

# Descargar archivo
curl -O https://ejemplo.com/archivo.zip

# Guardar con nombre específico
curl -o mi_archivo.zip https://ejemplo.com/archivo.zip

# Seguir redirecciones
curl -L https://ejemplo.com
```

---

### `ping`

Verifica conectividad de red.

**Sintaxis:**

```bash
ping [opciones] host
```

**Ejemplos:**

```bash
# Ping continuo (Ctrl+C para detener)
ping google.com

# Ping 5 veces
ping -c 5 google.com
```

---

## 12. Historial y Atajos

### `history`

Muestra el historial de comandos ejecutados.

**Ejemplos:**

```bash
# Ver historial
history

# Ver últimos 20 comandos
history 20

# Ejecutar comando por número
!123

# Ejecutar último comando
!!

# Ejecutar último comando que empezó con 'git'
!git

# Buscar en historial (Ctrl+R)
```

---

### Atajos de Teclado Útiles

**Navegación:**

- `Ctrl + A` - Ir al inicio de la línea
- `Ctrl + E` - Ir al final de la línea
- `Ctrl + U` - Borrar desde cursor hasta inicio
- `Ctrl + K` - Borrar desde cursor hasta final
- `Ctrl + W` - Borrar palabra anterior
- `Alt + B` - Retroceder una palabra
- `Alt + F` - Avanzar una palabra

**Control:**

- `Ctrl + C` - Cancelar comando actual
- `Ctrl + D` - Cerrar terminal (EOF)
- `Ctrl + Z` - Suspender proceso
- `Ctrl + L` - Limpiar pantalla (igual que `clear`)
- `Ctrl + R` - Buscar en historial

**Autocompletado:**

- `Tab` - Autocompletar comando/archivo
- `Tab Tab` - Mostrar opciones disponibles

---

## 13. Comandos de Ayuda

### `man` - Manual

Muestra el manual de cualquier comando.

**Sintaxis:**

```bash
man comando
```

**Navegación en man:**

- `Espacio` - Siguiente página
- `b` - Página anterior
- `/texto` - Buscar
- `q` - Salir

**Ejemplos:**

```bash
# Ver manual de ls
man ls

# Ver manual de grep
man grep
```

---

### `--help`

Muestra ayuda breve del comando.

**Ejemplos:**

```bash
ls --help
grep --help
tar --help
```

---

### `apropos`

Busca comandos por descripción.

**Ejemplo:**

```bash
# Buscar comandos relacionados con "copy"
apropos copy
```

---

## 14. Editores de Texto en Terminal

### `nano`

Editor de texto simple y amigable.

**Uso básico:**

```bash
nano archivo.txt
```

**Atajos en nano:**

- `Ctrl + O` - Guardar
- `Ctrl + X` - Salir
- `Ctrl + K` - Cortar línea
- `Ctrl + U` - Pegar
- `Ctrl + W` - Buscar
- `Ctrl + G` - Ayuda

---

### `vim` / `vi`

Editor de texto poderoso pero con curva de aprendizaje.

**Modos principales:**

- Modo Normal (navegación)
- Modo Inserción (edición)
- Modo Comando

**Comandos esenciales:**

```bash
# Abrir archivo
vim archivo.txt

# En modo normal:
i - Entrar modo inserción
Esc - Volver a modo normal
:w - Guardar
:q - Salir
:wq - Guardar y salir
:q! - Salir sin guardar
dd - Borrar línea
yy - Copiar línea
p - Pegar
```

---

## 15. Variables y Entorno

### Variables de Entorno

**Ver variables:**

```bash
# Mostrar todas las variables
env

# Mostrar variable específica
echo $HOME
echo $PATH
echo $USER
```

**Definir variables:**

```bash
# Variable temporal (solo en sesión actual)
VARIABLE="valor"

# Usar variable
echo $VARIABLE

# Variable permanente (añadir a ~/.bashrc)
export VARIABLE="valor"
```

**Variables comunes:**

- `$HOME` - Directorio home del usuario
- `$PATH` - Rutas donde buscar ejecutables
- `$USER` - Nombre del usuario actual
- `$PWD` - Directorio actual
- `$SHELL` - Shell actual

---

## 16. Consejos y Mejores Prácticas

### Seguridad

1. **No uses `sudo` innecesariamente** - Solo cuando realmente lo necesites
2. **Verifica antes de ejecutar `rm -rf`** - Esta operación es irreversible
3. **Lee scripts antes de ejecutarlos** - Especialmente los de Internet
4. **Usa contraseñas fuertes** - Para cuentas con privilegios
5. **Mantén backups** - De archivos importantes

### Productividad

1. **Usa Tab para autocompletar** - Ahorra tiempo y evita errores
2. **Aprende atajos de teclado** - Ctrl+R, Ctrl+A, Ctrl+E
3. **Crea alias para comandos frecuentes:**
   ```bash
   alias ll='ls -lah'
   alias ..='cd ..'
   alias update='sudo apt update && sudo apt upgrade'
   ```
4. **Usa historial** - Ctrl+R para buscar comandos anteriores
5. **Lee man pages** - La mejor documentación está incluida

### Organización

1. **Estructura de directorios clara** - Organiza tus archivos lógicamente
2. **Nombres descriptivos** - Para archivos y directorios
3. **Evita espacios en nombres** - Usa guiones o guiones bajos
4. **Documenta scripts** - Con comentarios (#)
5. **Versiona código importante** - Usa git

---

## 17. Comandos Avanzados (Introducción)

### `awk`

Procesamiento de texto basado en patrones.

**Ejemplo simple:**

```bash
# Imprimir segunda columna
ls -l | awk '{print $2}'

# Suma de números en columna
awk '{sum += $1} END {print sum}' numeros.txt
```

---

### `sed` - Stream Editor

Editor de flujo para transformar texto.

**Ejemplos:**

```bash
# Reemplazar texto
sed 's/viejo/nuevo/g' archivo.txt

# Eliminar líneas vacías
sed '/^$/d' archivo.txt

# Eliminar línea 5
sed '5d' archivo.txt
```

---

### `xargs`

Construye y ejecuta comandos desde entrada estándar.

**Ejemplos:**

```bash
# Eliminar archivos encontrados
find . -name "*.tmp" | xargs rm

# Crear directorios desde lista
cat lista.txt | xargs mkdir
```

---

## 18. Ejemplos Prácticos Combinados

### Encontrar archivos grandes

```bash
find . -type f -size +100M -exec ls -lh {} \; | sort -k 5 -rh
```

### Backup de directorio

```bash
tar -czvf backup_$(date +%Y%m%d).tar.gz /ruta/directorio/
```

### Buscar texto en múltiples archivos

```bash
grep -rn "función especial" /home/usuario/proyecto/
```

### Contar tipos de archivos

```bash
find . -type f | sed 's/.*\.//' | sort | uniq -c | sort -rn
```

### Monitorear uso de disco en tiempo real

```bash
watch -n 5 'df -h'
```

### Ver los 10 procesos que más memoria usan

```bash
ps aux | sort -rk 4 | head -10
```

### Encontrar y eliminar archivos duplicados

```bash
find . -type f -exec md5sum {} \; | sort | uniq -d -w32
```

---

## Recursos Adicionales

### Sitios web útiles

- https://explainshell.com - Explica comandos de shell
- https://tldr.sh - Resúmenes concisos de comandos
- https://ss64.com - Referencia de comandos

### Práctica

La mejor manera de aprender es **practicando**. No tengas miedo de experimentar (pero ten cuidado con `rm` y `sudo`).

---

## Glosario Rápido

- **Shell**: Intérprete de comandos (bash, zsh, etc.)
- **Terminal**: Aplicación que ejecuta el shell
- **Directorio**: Carpeta
- **Ruta absoluta**: Desde raíz `/home/usuario/docs`
- **Ruta relativa**: Desde ubicación actual `../docs`
- **Root**: Usuario administrador o directorio raíz `/`
- **Home**: Directorio personal `~` o `/home/usuario`
- **Pipe**: Tubería `|` que conecta comandos
- **Redirección**: `>` `>>` para guardar salidas
- **Wildcard**: Comodines `*` `?` para patrones
- **Flag/Opción**: Modificadores de comandos `-l` `--help`
