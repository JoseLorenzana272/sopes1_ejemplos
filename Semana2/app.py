import os
import socket
from flask import Flask, render_template
from redis import Redis

app = Flask(__name__)
# Conectamos a Redis (nombre del servicio en compose)
redis = Redis(host='redis', port=6379)

@app.route('/')
def hello():
    # Incrementamos contador
    count = redis.incr('hits')
    
    # Obtenemos variables para mostrar (Id del contenedor y color)
    hostname = socket.gethostname()
    color = os.getenv('COLOR_FONDO', '#2c3e50') # Color default gris oscuro
    
    # Renderizamos el HTML enviando los datos
    return render_template('index.html', visitas=count, hostname=hostname, color_fondo=color)

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True)