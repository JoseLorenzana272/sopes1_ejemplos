### 1. **Prepara tu entorno de desarrollo:**

- Asegúrate de tener instaladas las herramientas de compilación (`make`, `gcc`). En Ubuntu y sistemas basados en Debian, puedes instalarlas con:

  ```bash
  sudo apt-get update
  sudo apt-get install build-essential
  ```

- Instalar headers del kernel (`kernel headers`). Puedes instalarlos con:

  ```bash
  sudo apt-get install linux-headers-$(uname -r)
  ```

### 2. **Crea el archivo del módulo:**

- Abre un editor de texto y copia el código proporcionado en un archivo con extensión `.c`, por ejemplo, `modulo_ejemplo.c`.

### 3. **Crea un Makefile para compilar el módulo:**

- En el mismo directorio donde guardaste `modulo_ejemplo.c`, crea un archivo llamado `Makefile` con el siguiente contenido:

  ```makefile
  obj-m += modulo_ejemplo.o

  all:
     make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules

  clean:
     make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
  ```

- Este `Makefile` compila el módulo para tu versión actual del kernel.

### 4. **Compila el módulo:**

- En la terminal, navega al directorio donde guardaste `modulo_ejemplo.c` y `Makefile`. Luego ejecuta:

  ```bash
  make
  ```

- Si todo está correcto, se generará un archivo con la extensión `.ko` (por ejemplo, `modulo_ejemplo.ko`), que es tu módulo de kernel.

### 5. **Carga el módulo en el kernel:**

- Usa `insmod` para cargar el módulo:

  ```bash
  sudo insmod modulo_ejemplo.ko
  ```

- Para verificar que el módulo se ha cargado, usa:

  ```bash
  lsmod | grep modulo_ejemplo
  ```

- También puedes filtrar para buscar lo relacionado al modulo:

  ```bash
  sudo dmesg | grep -i "modulo"
  ```

- Deberías ver el JSON con la información del sistema.

### 6. **Desinstala el módulo cuando termines:**

- Para remover el módulo del kernel, usa:

  ```bash
  sudo rmmod modulo_ejemplo
  ```

- Verifica que el módulo fue removido correctamente:

  ```bash
  lsmod | grep modulo_ejemplo
  ```

### 7. **Depura y ajusta según sea necesario:**

- Si encuentras errores, puedes revisar los mensajes del kernel usando:

  ```bash
  dmesg | tail
  ```

- Ajusta el código del módulo para satisfacer todos los requisitos y, si es necesario, recompila el módulo.
