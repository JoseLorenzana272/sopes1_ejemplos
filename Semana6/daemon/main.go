package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/redis/go-redis/v9"
)

// Define los atributos principales de consumo simulado para un contenedor
type Container struct {
	Name   string
	CPU    float64
	Memory float64
	Status string
}

var ctx = context.Background()
var rdb *redis.Client

// Punto de entrada principal del daemon
func main() {
	// Verifica si el programa fue invocado por el cronjob
	if len(os.Args) > 1 && os.Args[1] == "crear-contenedor" {
		crearContenedorAleatorio()
		return
	}

	log.Println(("Starting daemon..."))
	// Inicializa los subsistemas necesarios
	conectarValkey()
	crearCronjob()

	// Inicia el bucle infinito para inyectar datos periodicamente
	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-ticker.C:
			// Simula las metricas y las guarda en memoria
			containers := generateRandomContainers()
			for _, c := range containers {
				key := fmt.Sprintf("container:%s:%d", c.Name, time.Now().Unix())
				data := map[string]interface{}{
					"name":   c.Name,
					"cpu":    c.CPU,
					"memory": c.Memory,
					"status": c.Status,
					"time":   time.Now().UnixMilli(),
				}
				err := rdb.HSet(ctx, key, data).Err()

				// Guardamos las métricas en un Stream para poder graficarlas fácilmente en Grafana
				streamKey := "stream:" + c.Name
				errStream := rdb.XAdd(ctx, &redis.XAddArgs{
					Stream: streamKey,
					Values: data,
				}).Err()

				// Almacena el valor mas reciente como clave simple para lectura directa de Grafana
				rdb.Set(ctx, "cpu:"+c.Name, fmt.Sprintf("%.2f", c.CPU), 2*time.Minute)
				rdb.Set(ctx, "memory:"+c.Name, fmt.Sprintf("%.2f", c.Memory), 2*time.Minute)
				rdb.Set(ctx, "status:"+c.Name, c.Status, 2*time.Minute)

				if err != nil || errStream != nil {
					log.Printf("Error saving container data: %v | %v", err, errStream)
				} else {
					rdb.Expire(ctx, key, 10*time.Minute)
					rdb.Expire(ctx, streamKey, 10*time.Minute)
					log.Printf("Saved container data: %s CPU %.2f Mem %.2f", c.Name, c.CPU, c.Memory)
				}
			}
		}
	}
}

// Establece la conexion de base de datos local contra el contenedor Valkey
func conectarValkey() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Valkey: %v", err)
	}

	log.Println("Connected to Valkey")
}

// Inyecta en el sistema operativo una instruccion cron para automatizar la ejecucion
func crearCronjob() {
	// Obtiene la ruta del ejecutable actual
	executable, _ := os.Executable()

	// Configura la expresion cron para ejecutar la tarea cada minuto (* * * * *)
	cronCommand := fmt.Sprintf("* * * * * %s crear-contenedor", executable)

	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(`(crontab -l 2>/dev/null; echo "%s") | crontab -`, cronCommand),
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error creating cronjob: %v, Output: %s", err, string(out))
	} else {
		log.Println("Cronjob created successfully")
	}
}

// Genera y arranca un contenedor Docker directamente sobre el sistema host
func crearContenedorAleatorio() {
	images := []string{"hello-world", "alpine"}
	randomImage := images[rand.Intn(len(images))]
	name := fmt.Sprintf("container_%d", time.Now().Unix())

	var cmd *exec.Cmd
	if randomImage == "alpine" {
		cmd = exec.Command("docker", "run", "-d", "--name", name, randomImage, "sleep", "300")
	} else {
		cmd = exec.Command("docker", "run", "-d", "--name", name, randomImage)
	}
	cmd.Run()
	log.Println("Created container: ", name)
}

// Produce metricas ficticias para un conjunto de contenedores y asi probar el monitoreo
func generateRandomContainers() []Container {
	names := []string{"nginx", "redis", "mysql", "postgres", "node"}
	status := []string{"running", "stopped"}

	// Define la cantidad aleatoria de contenedores a simular (entre 1 y 4)
	n := rand.Intn(4) + 1
	var containers []Container
	for i := 0; i < n; i++ {
		containers = append(containers, Container{
			Name:   names[rand.Intn(len(names))],
			CPU:    rand.Float64() * 100,
			Memory: rand.Float64() * 512,
			Status: status[rand.Intn(len(status))],
		})
	}
	return containers
}
