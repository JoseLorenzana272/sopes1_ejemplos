

# Asignar un nombre de host configurando `/etc/hosts`

## Pasos:

1. En las VMs, edita el archivo `/etc/hosts`:  
   ```bash
   sudo nano /etc/hosts
   ```

2. Agrega las IPs de las VMs con un nombre personalizado:  
   ```
   192.168.100.10 vm1
   192.168.100.11 vm2
   192.168.122.14 vmdocker
   ```

3. En el codigo Go de las APIs, puedes llamar a cada api ahora por el nombre de la vm donde estan:  
   ```go
   http.Get("http://vm1:8081/")
   http.Get("http://vm2:8082/")
   ```

