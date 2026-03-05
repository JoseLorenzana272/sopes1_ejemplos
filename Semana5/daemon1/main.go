package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Abrir o crear un archivo log
	f, err := os.OpenFile("/var/log/mydaemon.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // rw-r--r--
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// escribir en el archivo de logs
	log.SetOutput(f)
	log.Println("Daemon started, let's read the Memory :D")

	// loop infinito
	for {
		// leer memoria desde /proc/meminfo
		total, free, used := readMemoryFromProc()

		// imprimir metricas
		log.Printf("Total RAM Memory: %d kB, Free Memory: %d kB, Used Memory: %d kB\n", total, free, used)

		time.Sleep(5 * time.Second)

	}
}

// leer memoria total y libre desde y usada /proc/meminfo
func readMemoryFromProc() (totalKB, freeKB, usedKB uint64) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Printf("Error opening /proc/meminfo: %v", err)
		return 0, 0, 0
	}
	defer file.Close()

	// leer lineapor linea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		field := strings.Fields(line)
		if len(field) < 2 {
			continue
		}

		value, err := strconv.ParseUint(field[1], 10, 64)
		if err != nil {
			continue
		}

		switch field[0] {
		case "MemTotal:":
			totalKB = value
		case "MemFree:":
			freeKB = value
		}
	}

	// RAm usada = RAM total - RAM libre
	if totalKB > freeKB {
		usedKB = totalKB - freeKB
	}

	return
}
