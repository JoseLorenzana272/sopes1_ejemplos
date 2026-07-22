# Comandos Básicos de Linux

## Navegación y Gestión de Archivos

### Navegación en el Sistema de Archivos

```bash
# Mostrar directorio actual
pwd

# Listar contenido del directorio
ls                  # Lista básica
ls -l              # Lista detallada
ls -la             # Lista detallada incluyendo archivos ocultos
ls -lh             # Lista con tamaños legibles por humanos

# Cambiar de directorio
cd /home/usuario   # Ir a directorio específico
cd ..              # Subir un nivel
cd ~               # Ir al directorio home
cd -               # Volver al directorio anterior
```

### Gestión de Archivos y Directorios

```bash
# Crear directorios
mkdir mi_directorio
mkdir -p ruta/completa/nueva    # Crear toda la estructura

# Crear archivos
touch archivo.txt
echo "Contenido" > archivo.txt  # Crear archivo con contenido
echo "Más contenido" >> archivo.txt  # Agregar contenido

# Copiar archivos y directorios
cp archivo.txt copia.txt
cp -r directorio/ copia_directorio/  # Copiar recursivamente

# Mover/renombrar archivos
mv archivo.txt nuevo_nombre.txt
mv archivo.txt /otra/ubicacion/

# Eliminar archivos y directorios
rm archivo.txt
rm -f archivo.txt              # Forzar eliminación
rm -r directorio/              # Eliminar recursivamente
rm -rf directorio/             # Eliminar recursivamente sin confirmar
```

### Visualización de Contenido

```bash
# Mostrar contenido de archivos
cat archivo.txt                # Mostrar todo el contenido
less archivo.txt               # Ver archivo página por página
head archivo.txt               # Primeras 10 líneas
head -n 20 archivo.txt         # Primeras 20 líneas
tail archivo.txt               # Últimas 10 líneas
tail -f /var/log/syslog        # Seguir archivo en tiempo real
```

## Búsqueda y Filtrado

```bash
# Buscar archivos
find /home -name "*.txt"       # Buscar archivos .txt en /home
find . -type f -name "config*" # Buscar archivos que empiecen con "config"
locate archivo.txt             # Búsqueda rápida en base de datos

# Buscar dentro de archivos
grep "patrón" archivo.txt      # Buscar patrón en archivo
grep -r "patrón" directorio/   # Búsqueda recursiva
grep -i "patrón" archivo.txt   # Búsqueda sin distinguir mayúsculas
grep -n "patrón" archivo.txt   # Mostrar números de línea
```

## Información del Sistema

```bash
# Información del sistema
uname -a                       # Información del kernel
whoami                         # Usuario actual
id                             # ID de usuario y grupos
date                           # Fecha y hora actual
uptime                         # Tiempo de funcionamiento del sistema
df -h                          # Espacio en disco
du -sh directorio/             # Tamaño de directorio
free -h                        # Memoria RAM disponible
ps aux                         # Procesos en ejecución
top                            # Monitor de procesos en tiempo real
htop                           # Monitor mejorado (si está instalado)
```

## Gestión de Usuarios y Grupos

### Gestión de Usuarios

```bash
# Crear nuevo usuario
sudo useradd juan
sudo useradd -m -s /bin/bash juan  # Con directorio home y shell

# Establecer contraseña
sudo passwd juan

# Modificar usuario
sudo usermod -c "Juan Pérez" juan  # Cambiar nombre completo
sudo usermod -s /bin/zsh juan      # Cambiar shell
sudo usermod -g developers juan   # Cambiar grupo principal
sudo usermod -aG sudo juan        # Agregar a grupo sudo (administradores)

# Eliminar usuario
sudo userdel juan                  # Solo eliminar usuario
sudo userdel -r juan              # Eliminar usuario y su directorio home

# Ver información de usuarios
cat /etc/passwd                   # Lista de usuarios
finger juan                       # Información detallada del usuario
id juan                          # IDs y grupos del usuario
```

### Gestión de Grupos

```bash
# Crear grupo
sudo groupadd developers
sudo groupadd -g 1001 developers  # Con GID específico

# Agregar usuario a grupo
sudo usermod -aG developers juan
sudo gpasswd -a juan developers   # Método alternativo

# Eliminar usuario de grupo
sudo gpasswd -d juan developers
sudo deluser juan developers      # En sistemas Debian/Ubuntu

# Ver grupos
cat /etc/group                    # Lista de todos los grupos
groups juan                       # Grupos del usuario juan
getent group developers          # Información del grupo

# Eliminar grupo
sudo groupdel developers
```

### Cambio de Usuario

```bash
# Cambiar a otro usuario
su juan                          # Cambiar a usuario juan
su - juan                        # Cambiar con entorno completo
sudo -u juan comando            # Ejecutar comando como juan
sudo su -                       # Cambiar a root con entorno completo
```

## Sistema de Permisos

### Entendiendo los Permisos

Los permisos en Linux se representan con 3 grupos de 3 bits cada uno:
- **Usuario propietario (u)**: rwx
- **Grupo propietario (g)**: rwx  
- **Otros usuarios (o)**: rwx

Donde:
- **r** = read (lectura) = 4
- **w** = write (escritura) = 2
- **x** = execute (ejecución) = 1

### Visualizar Permisos

```bash
# Ver permisos detallados
ls -l archivo.txt
# Salida: -rw-r--r-- 1 usuario grupo 1024 jan 15 10:30 archivo.txt
#         ↑ permisos  propietario grupo

# Explicación de -rw-r--r--:
# - = tipo de archivo (- = archivo, d = directorio, l = enlace)
# rw- = propietario puede leer y escribir
# r-- = grupo puede solo leer  
# r-- = otros pueden solo leer
```

### Cambiar Permisos con chmod

```bash
# Método numérico (octal)
chmod 755 script.sh              # rwxr-xr-x
chmod 644 archivo.txt            # rw-r--r--
chmod 600 archivo_privado.txt    # rw-------
chmod 777 archivo.txt            # rwxrwxrwx (todos los permisos)

# Método simbólico
chmod u+x script.sh              # Agregar ejecución al propietario
chmod g+w archivo.txt            # Agregar escritura al grupo
chmod o-r archivo.txt            # Quitar lectura a otros
chmod a+r archivo.txt            # Agregar lectura a todos (all)
chmod u=rwx,g=rx,o=r archivo.txt # Establecer permisos específicos

# Cambiar permisos recursivamente
chmod -R 755 directorio/         # Aplicar a directorio y contenido
```

### Cambiar Propietario con chown

```bash
# Cambiar propietario
sudo chown juan archivo.txt

# Cambiar propietario y grupo
sudo chown juan:developers archivo.txt

# Solo cambiar grupo
sudo chown :developers archivo.txt
sudo chgrp developers archivo.txt    # Método alternativo

# Cambiar recursivamente
sudo chown -R juan:developers directorio/
```

### Ejemplos Prácticos de Permisos

```bash
# Crear script ejecutable
echo '#!/bin/bash' > mi_script.sh
echo 'echo "Hola Mundo"' >> mi_script.sh
chmod +x mi_script.sh
./mi_script.sh

# Archivo solo para el propietario
echo "Información confidencial" > secreto.txt
chmod 600 secreto.txt

# Directorio compartido para un grupo
sudo mkdir /compartido
sudo chown root:developers /compartido
sudo chmod 775 /compartido
sudo chmod g+s /compartido          # SetGID para herencia de grupo

# Ver permisos especiales
ls -la /tmp                         # Sticky bit en /tmp (t)
ls -la /usr/bin/passwd              # SetUID en passwd (s)
```

### Permisos Especiales

```bash
# SetUID (s en posición de x del propietario)
chmod u+s archivo                   # 4755
# El archivo se ejecuta con permisos del propietario

# SetGID (s en posición de x del grupo)  
chmod g+s directorio               # 2755
# Archivos creados heredan el grupo del directorio

# Sticky Bit (t en posición de x de otros)
chmod +t directorio                # 1755
# Solo el propietario puede eliminar sus archivos
```

## Comandos de Red y Conectividad

```bash
# Información de red
ip addr show                       # Mostrar interfaces de red
ip route show                      # Mostrar tabla de rutas
ping google.com                    # Probar conectividad
wget https://ejemplo.com/archivo   # Descargar archivo
curl -O https://ejemplo.com/archivo # Descargar con curl

# Información de puertos
netstat -tuln                      # Puertos abiertos
ss -tuln                          # Alternativa moderna a netstat
```

## Gestión de Procesos

```bash
# Ver procesos
ps aux                            # Todos los procesos
ps aux | grep apache             # Filtrar procesos
pgrep apache                     # PIDs de procesos apache
pkill apache                     # Terminar procesos apache

# Control de procesos
comando &                        # Ejecutar en segundo plano
jobs                            # Ver trabajos en segundo plano
fg %1                          # Traer trabajo 1 al primer plano
bg %1                          # Enviar trabajo 1 al segundo plano
nohup comando &                # Ejecutar independiente del terminal

# Terminar procesos
kill PID                       # Terminar proceso por PID
kill -9 PID                   # Forzar terminación
killall nombre_proceso        # Terminar por nombre
```

## Compresión y Archivos

```bash
# Archivos tar
tar -cvf archivo.tar directorio/  # Crear archivo tar
tar -xvf archivo.tar              # Extraer archivo tar
tar -czvf archivo.tar.gz dir/     # Crear tar comprimido con gzip
tar -xzvf archivo.tar.gz          # Extraer tar.gz

# Archivos zip
zip -r archivo.zip directorio/    # Crear archivo zip
unzip archivo.zip                 # Extraer archivo zip
```

## Variables de Entorno

```bash
# Ver variables de entorno
env                              # Todas las variables
echo $PATH                       # Variable PATH
echo $HOME                       # Directorio home
echo $USER                       # Usuario actual

# Establecer variables
export MI_VARIABLE="valor"       # Variable temporal
echo 'export MI_VAR="valor"' >> ~/.bashrc  # Variable permanente
```

Estos comandos cubren las operaciones más comunes en Linux, desde navegación básica hasta gestión avanzada de usuarios, grupos y permisos. Practica estos comandos en un entorno seguro para familiarizarte con su uso.