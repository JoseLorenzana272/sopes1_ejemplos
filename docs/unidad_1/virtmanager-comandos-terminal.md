
# Documentación: Uso de Virt-Manager para gestionar máquinas virtuales

Virt-Manager es una interfaz gráfica para administrar máquinas virtuales (VMs) usando libvirt. Permite crear, eliminar, visualizar, y controlar VMs de manera sencilla.

## Instalación

En la mayoría de distribuciones Linux:

```bash
sudo apt install virt-manager         # Debian/Ubuntu
sudo dnf install virt-manager         # Fedora
sudo pacman -S virt-manager           # Arch
```

## Abrir Virt-Manager

```bash
virt-manager
```
O busca "Virtual Machine Manager" en el menú de aplicaciones.

## Ver tus máquinas virtuales

Al abrir Virt-Manager, verás una lista de todas las VMs disponibles en tu sistema. Cada VM muestra su estado (corriendo, detenida, etc.).

## Ver la IP de una VM

1. Haz doble clic en la VM para abrir su ventana.
2. Ve al menú `Ver` > `Detalles`.
3. Selecciona la pestaña `Interfaces de red` o `NIC`.
4. Allí verás la IP asignada (si la VM está encendida y tiene red configurada).

**Alternativa:**
Puedes abrir la consola de la VM y usar comandos como `ip a` o `ifconfig` dentro del sistema operativo invitado.

## Otras funciones útiles

- **Crear una nueva VM:** Haz clic en el ícono de "Monitor con estrella" o en `Archivo > Nueva máquina virtual` y sigue el asistente.
- **Clonar una VM:** Haz clic derecho sobre la VM > `Clonar`.
- **Tomar snapshots:** Haz clic derecho > `Tomar instantánea` para guardar el estado actual de la VM.
- **Conectar dispositivos USB:** Ve a `Detalles` > `Agregar hardware` > `USB Host Device`.
- **Ajustar recursos:** En `Detalles` puedes modificar CPU, RAM, almacenamiento, etc. (la VM debe estar apagada).

## Comandos útiles con virsh (opcional)

Virt-Manager usa `libvirt` por debajo, así que puedes usar comandos como:

```bash
virsh list --all           # Lista todas las VMs
virsh domifaddr <nombre>   # Muestra la IP de una VM
virsh start <nombre>       # Inicia una VM
virsh shutdown <nombre>    # Apaga una VM
```

---
**Referencia:** [Virt-Manager Documentation](https://virt-manager.org/) 
