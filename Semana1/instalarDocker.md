## Instalar Docker - SOPES1

### 1. Actualizar sistema

```bash
sudo apt update
```


### 2. Docker con el paquete oficial


```bash
curl -fsSL https://get.docker.com | sudo sh
```

### 3. Prueba de que se instaló

```bash
docker --version
```

Probar solo para ver si está todo bien:

```bash
sudo docker run hello-world
```

## Opcional para usar Docker sin sudo


```bash
sudo usermod -aG docker $USER
```

Cerrar y volver a entrar para probar


```bash
docker ps
```

## Habilitar Docker al iniciar el todo sistema

esto es solo por si acaso, regularmente se pone solo

```bash
sudo systemctl enable docker
sudo systemctl start docker
```

:)