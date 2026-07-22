


# Solicitar al usuario la IP de la VM Docker, el nombre de la imagen y el nombre del contenedor
read -p "Ingrese la IP de la VM con Docker: " IP_VM_DOCKER
read -p "Ingrese el nombre de la imagen a descargar: " IMAGE_NAME
read -p "Ingrese el nombre del contenedor a crear: " CONTAINER_NAME

echo "Iniciando el contenedor '$CONTAINER_NAME' desde el registro privado..."

# Descargar la imagen desde el registro privado usando los datos ingresados
sudo ctr images pull --plain-http ${IP_VM_DOCKER}:5000/${IMAGE_NAME}

# Ejecutar el contenedor en segundo plano con red de host usando los datos ingresados
sudo ctr run -d --tty --net-host ${IP_VM_DOCKER}:5000/${IMAGE_NAME} ${CONTAINER_NAME}

echo "Contenedor '$CONTAINER_NAME' iniciado exitosamente en segundo plano."

# Listar todas las imágenes disponibles en containerd
sudo ctr images list

# Listar todas las tareas (contenedores en ejecución) en containerd
sudo ctr tasks list


# Para ejecutar este script, guárdalo como `archivo-de-ejecucion.sh` y dale permisos de ejecución:
# chmod +x archivo-de-ejecucion.sh
# Luego, puedes ejecutarlo con:
# ./archivo-de-ejecucion.sh

# NOTA: si no puedes conectarte a tu máquina desde VSCode para crear y pegar el contenido en el script, puedes crear el archivo directamente en la terminal con:
# cat > archivo-de-ejecucion.sh
# Luego pega el contenido del script, presiona enter y por último presiona Ctrl+D para guardar.
# Listo, ya creaste el script `archivo-de-ejecucion.sh` en tu máquina virtual, ahora puedes darle permisos de ejecución y ejecutarlo.

# NOTA2: puede que te pida tu contraseña de usuario para ejecutar los comandos con `sudo`. Si es así, ingresa tu contraseña de usuario cuando te la pida, sino ejecuta el archivo directamente con:
# sudo ./archivo-de-ejecucion.sh
# Esto es necesario para que los comandos que requieren permisos de superusuario se ejecuten correctamente.
