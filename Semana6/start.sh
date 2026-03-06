#!/bin/bash

echo "--------------------------------"
echo "Compilando daemon..."

if ! go build -o daemon-bin daemon/main.go; then
    echo "[ERROR] No se pudo compilar el daemon"
    exit 1
fi

echo "Daemon compilado correctamente"
echo "--------------------------------"

echo "Iniciando contenedores Grafana y Valkey..."

cd grafana-valkey || exit

docker compose up -d

cd ..

echo "--------------------------------"
echo "Iniciando daemon..."

./daemon-bin &

DAEMON_PID=$!

echo $DAEMON_PID > daemon.pid

echo "Daemon iniciado con PID $DAEMON_PID"

echo "--------------------------------"
echo "Sistema iniciado correctamente"

echo "Grafana disponible en:"
echo "http://localhost:3001"