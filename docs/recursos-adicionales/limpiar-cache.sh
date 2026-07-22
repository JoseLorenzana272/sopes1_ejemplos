#!/bin/bash

# Limpiar caché de APT (paquetes)
sudo apt clean

# Limpiar paquetes no necesarios
sudo apt autoremove --purge -y

# Limpiar caché de thumbnails
rm -rf ~/.cache/thumbnails/*

# Limpiar caché de aplicaciones (opcional)
rm -rf ~/.cache/*

# Limpiar archivos temporales
sudo rm -rf /tmp/*

# Limpiar caché de memoria RAM (opcional, se libera memoria inactiva)
sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches

echo "¡Limpieza completada!"

# Para poder ejecutar este script, guárdalo como `limpiar-cache.sh` y dale permisos de ejecución:
# chmod +x limpiar-cache.sh
# Luego, puedes ejecutarlo con:
# ./limpiar-cache.sh


# NOTA: si no puedes conectarte a tu maquina desde VSCode para crear y pegar el contenido en el script, puedes crear el archivo directamente en la terminal con:
# cat > limpiar-cache.sh
# Luego pega el contenido del script, presiona enter y por ultimo presiona Ctrl+D para guardar.
# Listo ya creaste el script `limpiar-cache.sh` en tu maquina virtual ahora puedes darle permisos de ejecucion y ejecutarlo.


# NOTA2: puede que te pida tu contraseña de usuario para ejecutar los comandos con `sudo` si es asi ingresa tu contraseña de usuario si te la pide sino ejecuta el archivo directamente con:
# sudo ./limpiar-cache.sh
# Esto es necesario para que los comandos que requieren permisos de superusuario se ejecuten correctamente.

