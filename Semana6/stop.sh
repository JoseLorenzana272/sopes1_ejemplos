#!/bin/bash

echo "--------------------------------"
echo "Deteniendo daemon..."

if [ -f daemon.pid ]; then
    PID=$(cat daemon.pid)
    kill $PID
    rm daemon.pid
    echo "Daemon detenido"
else
    echo "No se encontró daemon.pid"
fi

echo "--------------------------------"
echo "Eliminando cronjob..."

crontab -r 2>/dev/null

echo "Cronjob eliminado"

echo "--------------------------------"
echo "Deteniendo docker-compose..."

cd grafana-valkey || exit

docker compose down

cd ..

echo "--------------------------------"
echo "Sistema detenido correctamente"