
# Instrucciones para usar Zot como registro de contenedores

## ¿Qué es Zot?

**Zot** es un registro de contenedores de código abierto y liviano que implementa la especificación OCI (Open Container Initiative) Distribution. Es una alternativa moderna y eficiente a otros registros como Docker Registry, diseñado para ser:

- **Ligero y rápido**: Consume pocos recursos del sistema
- **Seguro**: Incluye características de seguridad avanzadas como verificación de firmas
- **Compatible con OCI**: Funciona con cualquier cliente compatible con Docker/OCI
- **Fácil de configurar**: Configuración mínima requerida para funcionar

Zot es ideal para desarrollo local, CI/CD pipelines y despliegues en entornos donde necesitas un registro privado de contenedores.

---

# Comando para usar ZOT

1. **Iniciar el registro Zot en tu VM con DOCKER**

    Ejecuta el siguiente comando para iniciar un registro Zot en segundo plano, exponiendo el puerto 5000:
    ```sh
    docker run -d -p 5000:5000 --name zot ghcr.io/project-zot/zot-linux-amd64:latest
    ```
    Esto descargará la imagen de Zot y la ejecutará como un contenedor llamado `zot`.

    
## En tus Computadora donde Desarrollaste todo:
2. **Edita la configuración de Docker para que pueda subir la imagen a la VM con DOCKER**

   a. **Editar la configuración de Docker**:
    ```bash
    sudo nano /etc/docker/daemon.json
    ```

   b. **Agregar o modificar el contenido del archivo** (si está vacío, agregar lo siguiente):
    ```json
    {
        "insecure-registries": ["<IP_VM_DOCKER>:5000"]
    }
    ```

   c. **Reiniciar Docker** para aplicar los cambios:
    ```bash
    sudo systemctl restart docker
    ```

3. **Etiquetar la imagen para el registro privado**

    Usaremos de ejemplo la imagen llamada fiber-api-go.
    Cambia la etiqueta de la imagen para que apunte a tu registro privado (reemplaza `<IP_VM_DOCKER>` por la IP de tu VM):
    ```sh
    docker tag fiber-api-go:v1 <IP_VM_DOCKER>:5000/fiber-api-go:v1
    ```

4. **Subir la imagen al registro Zot**

    Sube la imagen etiquetada a tu registro Zot:
    ```sh
    docker push <IP_VM_DOCKER>:5000/fiber-api-go:v1
    ```

5. **Verificar las imágenes disponibles en el registro**

    Consulta el catálogo de imágenes almacenadas en el registro Zot:
    ```sh
    curl http://<IP_VM_DOCKER>:5000/v2/_catalog
    ```
   ***También puedes pegar la URL en tu navegador para verificar que funciona***

## En tus maquinas virtuales con unicamente containerd y ctr
6. **Descargar la imagen desde el registro Zot**

    Descarga la imagen desde tu registro privado para comprobar que está disponible:
    ```sh
    sudo ctr images pull --plain-http <IP_VM_DOCKER>:5000/fiber-api-go:v1
    ```
    Lista las imagenes
    ```sh
    sudo ctr images ls

    # si el comando anterior no funciona prueba con este
    sudo ctr images list
    ```

    Listo ya puedes conectarte con tu Registro de Contenedores privados de ZOT en tu maquina virtual, ahora puedes dirigirste a ``containerd-instrucciones.md`` para ver los comandos basicos de ctr (ctr = cliente de consola de containerd) para que puedas ejecutar tus contenedores con tus imagenes propias.


