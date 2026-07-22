# Instrucciones y Scripts para Gestión de Contenedores

A continuación, se presenta un índice con una breve descripción de cada archivo:

## Índice de Archivos

- [instalar-containerd.sh](./instalar-containerd.sh):
  Script para instalar containerd en sistemas basados en Ubuntu, incluyendo la configuración de repositorios y dependencias necesarias.

- [archivo-de-ejecucion.sh](./archivo-de-ejecucion.sh):
  Script interactivo para descargar una imagen desde un registro privado y ejecutar un contenedor usando containerd (`ctr`). Solicita la IP del registro al usuario y automatiza el proceso.

- [limpiar-cache.sh](./limpiar-cache.sh):
  Script para limpiar la caché del sistema, paquetes innecesarios y archivos temporales, ayudando a mantener el sistema optimizado.

- [zot-intrucciones.md](./zot-intrucciones.md):
  Guía paso a paso para instalar y utilizar Zot como registro privado de contenedores, incluyendo la configuración de Docker y el manejo de imágenes.

- [containerd-instrucciones.md](./containerd-instrucciones.md):
  Manual de comandos básicos para el uso de `ctr` (cliente de containerd), desde la gestión de imágenes hasta la ejecución y administración de contenedores.

---

Cada archivo está diseñado para facilitar la administración de contenedores y la integración con registros privados. Para más detalles, consulta cada documento o script individualmente.
