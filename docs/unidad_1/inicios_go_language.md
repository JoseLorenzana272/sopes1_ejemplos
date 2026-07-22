# Introducción a Go

Go, también conocido como Golang, es un lenguaje de programación desarrollado por Google. Es conocido por su simplicidad, eficiencia y soporte para la concurrencia. A continuación, se presentan algunos conceptos básicos para entender cómo funciona Go:

## Características principales
- **Compilado y estáticamente tipado**: Go es un lenguaje compilado, lo que significa que el código se traduce a un binario ejecutable antes de ejecutarse. Además, es estáticamente tipado, lo que significa que los tipos de datos se verifican en tiempo de compilación.
- **Concurrencia**: Go tiene soporte nativo para la concurrencia a través de goroutines y canales.
- **Recolección de basura**: Go incluye un recolector de basura para manejar la memoria automáticamente.
- **Sintaxis simple**: La sintaxis de Go es fácil de aprender y leer.

## Primer programa en Go
A continuación, se muestra un ejemplo de un programa básico en Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, mundo!")
}
```

### Explicación del código
1. `package main`: Define el paquete principal del programa.
2. `import "fmt"`: Importa el paquete `fmt` para imprimir mensajes en la consola.
3. `func main()`: Define la función principal donde comienza la ejecución del programa.
4. `fmt.Println("¡Hola, mundo!")`: Imprime el mensaje "¡Hola, mundo!" en la consola.

## Instalación de Go
1. Descarga Go desde su [sitio oficial](https://golang.org/dl/).
2. Sigue las instrucciones de instalación para tu sistema operativo.
3. Verifica la instalación ejecutando `go version` en la terminal.

## Compilación y ejecución
- Para compilar un programa en Go, usa el comando `go build nombre_archivo.go`.
- Para ejecutarlo directamente sin compilar, usa `go run nombre_archivo.go`.

## Recursos adicionales
- [Documentación oficial de Go](https://golang.org/doc/)
- [Tour de Go](https://tour.golang.org/)