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

type Container struct {
	Name   string
	CPU    float64
	Memory float64
	Status string
}

var ctx = context.Background()
var rdb *redis.Client

func main() {
	if len(os.Args) > 1 && os.Args[1] == "crear-contenedor" {
		crearContenedorAleatorio()
		return
	}

	log.Println(("Starting daemon..."))
	conectarValkey()
	crearCronjob()

	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-ticker.C:
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
				if err != nil {
					log.Printf("Error saving container data: %v", err)
				} else {
					rdb.Expire(ctx, key, 10*time.Minute)
					log.Printf("Saved container data: %s CPU %.2f Mem %.2f", c.Name, c.CPU, c.Memory)
				}
			}
		}
	}
}

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

func crearCronjob() {
	// obtener el path del ejecutable actual
	executable, _ := os.Executable()

	// que se ejecute cada minuti * * * * *
	cronCommand := fmt.Sprintf("* * * * * %s crear-contenedor", executable)

	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(`(crontab -l 2>/dev/null; echo "%s") | crontab -`, cronCommand),
	)
	cmd.Run()
	log.Println("Cronjob created")
}

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

func generateRandomContainers() []Container {
	names := []string{"nginx", "redis", "mysql", "postgres", "node"}
	status := []string{"running", "stopped"}

	// generar entre 1 y 4 contenedores
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
