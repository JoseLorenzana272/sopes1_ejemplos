echo "Instalando containerd..."

# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

# Install containerd from the Docker repository:
sudo apt-get install containerd.io

echo "Containerd instalado exitosamente!"


# NOTA: Si no puedes conectarte a tu máquina desde VSCode para crear y pegar el contenido en el script, puedes crear el archivo directamente en la terminal con:
# cat > instalar-containerd.sh
# Luego pega el contenido del script, presiona enter y por último presiona Ctrl+D para guardar.
# Listo, ya creaste el script `instalar-containerd.sh` en tu máquina virtual. Ahora puedes darle permisos de ejecución y ejecutarlo.
# chmod +x instalar-containerd.sh
# Luego, puedes ejecutarlo con:
# ./instalar-containerd.sh


# NOTA2: puede que te pida tu contraseña de usuario para ejecutar los comandos con `sudo` si es asi ingresa tu contraseña de usuario si te la pide sino ejecuta el archivo directamente con:
# sudo ./instalar-containerd.sh
# Esto es necesario para que los comandos que requieren permisos de superusuario se ejecuten correctamente.