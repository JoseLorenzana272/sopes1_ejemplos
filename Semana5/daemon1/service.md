[Unit]
Description=Mi daemon en Go
After=network.target

[Service]
# Ruta donde se encuentra el binario
ExecStart=/usr/local/bin/mydaemon
Restart=always

[Install]
WantedBy=multi-user.target
