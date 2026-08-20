# Ejemplo de Daemon en Go

## Requisitos previos

- Go instalado en el sistema. En caso de usar Ubuntu Server, puedes instalarlo ejecutando los siguientes comandos:
  ```bash
  sudo apt update
  sudo apt install golang-go -y
  ```

## Acerca del archivo service.md

El archivo `service.md` contiene la definicion de la unidad de systemd necesaria para que el sistema operativo maneje nuestro programa compilado en Go como un servicio en segundo plano (daemon). Al copiar este archivo a `/etc/systemd/system/`, el sistema lee esta configuracion para saber donde encontrar el ejecutable, cuando debe arrancarlo y bajo que condiciones debe reiniciarse.

## Instrucciones de ejecución

### 1. Compilar el programa

Navega al directorio donde se encuentra el código fuente y compila el programa en Go:

```bash
cd daemon1
go build -o mydaemon main.go
```

### 2. Mover el binario

Mueve el archivo binario generado al directorio ejecutable del sistema:

```bash
sudo mv mydaemon /usr/local/bin/
```

### 3. Configurar el servicio systemd

Copia el archivo de configuración proporcionado (`service.md`) al directorio de systemd:

```bash
sudo cp daemon1/service.md /etc/systemd/system/mydaemon.service
```

### 4. Recargar el demonio de systemd

Recarga la configuración de systemd para que reconozca el nuevo servicio:

```bash
sudo systemctl daemon-reload
```

### 5. Iniciar y habilitar el servicio

Inicia el servicio y habilítalo para que se ejecute automáticamente al arrancar el sistema:

```bash
sudo systemctl start mydaemon
sudo systemctl enable mydaemon
```

### 6. Verificar el estado del servicio

Comprueba que el servicio se está ejecutando correctamente:

```bash
sudo systemctl status mydaemon
```

### 7. Revisar los registros (logs)

El daemon escribe métricas de uso de memoria en `/var/log/mydaemon.log`. Para visualizar los registros en tiempo real, ejecuta:

```bash
sudo tail -f /var/log/mydaemon.log
```

## Detener y eliminar el servicio

Si deseas detener y deshabilitar el servicio, ejecuta los siguientes comandos:

```bash
sudo systemctl stop mydaemon
sudo systemctl disable mydaemon
sudo rm /etc/systemd/system/mydaemon.service
sudo systemctl daemon-reload
sudo rm /usr/local/bin/mydaemon
```
