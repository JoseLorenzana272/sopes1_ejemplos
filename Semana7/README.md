# API de Clima en Rust

API REST simple desarrollada en Rust utilizando el framework `actix-web`. La API permite consultar un mensaje de bienvenida y enviar datos meteorológicos para recibir un resumen formateado.

## Endpoints Disponibles

La API cuenta con los siguientes endpoints:

### 1. GET `/`

Devuelve un mensaje de texto de bienvenida.

**Respuesta:**

```text
Hi, welcome to Climate API
```

### 2. POST `/clima`

Recibe datos sobre el clima en formato JSON y devuelve una oración descriptiva basada en esos datos.

**Cuerpo de la petición (JSON):**

```json
{
  "ciudad": "Guatemala",
  "temperatura": 25.5,
  "humedad": 60.0,
  "condicion": "Soleado"
}
```

**Respuesta (JSON):**

```json
{
  "mensaje": "El clima en Guatemala es de 25.5°C con una humedad del 60% y condición de Soleado.",
  "datos_recibidos": {
    "ciudad": "Guatemala",
    "temperatura": 25.5,
    "humedad": 60.0,
    "condicion": "Soleado"
  }
}
```

## Requisitos

Para compilar y ejecutar el proyecto, es necesario tener instalado Rust y su gestor de paquetes `cargo`. Puede instalar las herramientas necesarias desde [rustup.rs](https://rustup.rs/).

## Instrucciones de Ejecución

1. Abra una terminal y navegue hacia el directorio del proyecto:

   ```bash
   cd api-rust
   ```

2. Ejecute el servidor utilizando el comando `cargo`:

   ```bash
   cargo run
   ```

   La primera vez que se ejecute el comando, se descargarán y compilarán todas las dependencias definidas.

3. El servidor se iniciará y escuchará peticiones en la siguiente dirección:

   ```text
   http://127.0.0.1:8080
   ```

4. Puede probar el funcionamiento de la API utilizando una herramienta como `curl`:
   ```bash
   curl -X POST http://127.0.0.1:8080/clima \
        -H "Content-Type: application/json" \
        -d '{"ciudad": "Guatemala", "temperatura": 25.5, "humedad": 60.0, "condicion": "Soleado"}'
   ```
