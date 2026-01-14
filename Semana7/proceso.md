
### Preparación 
 **dos terminales** abiertas una al lado de la otra.

* **Terminal 1:** Será la víctima (El proceso).
* **Terminal 2:** Será el investigador (Tú).

---

### Parte 1: Creando al "Paciente"

En la **Terminal 1**, vamos a crear un proceso que no haga nada más que existir, para poder estudiarlo. Usaremos Python (o un simple `sleep`) para no gastar recursos.

```bash
# Este comando imprime su PID y se duerme por 1 hora
python3 -c "import os, time; print(f'>>> MI PID ES: {os.getpid()} <<<'); time.sleep(3600)"

```

---

### Parte 2: La visión "Usuario" vs "Kernel"

En la **Terminal 2**, muestra la diferencia.

1. **Visión de Usuario (Comando `ps`):**
```bash
ps -fp 1234  # (Reemplazar 1234 con el PID real)

```


* *Explicación:* "Esto es lo que ve el usuario. Un resumen bonito."


2. **Visión del Kernel (Directo al `task_struct`):**
Explica que Linux guarda cada proceso en una carpeta con su número.
```bash
cd /proc/1234
ls

```


* *Wow factor:* Diles: *"Miren todos estos archivos. Cada uno representa un campo de la estructura `task_struct` que vimos en las diapositivas."*



---

### Parte 3: Mapeando la Teoría a la Realidad

Ahora usa `cat` para leer los archivos y conéctalos con tu PDF de la Semana 7:

**A. El Estado del Proceso (`state`)**
En tu PDF se habla de `TASK_RUNNING`, `TASK_INTERRUPTIBLE`, etc.

```bash
cat status | grep State

```

* **Resultado:** Verás `S (sleeping)`.
* **Lección:** "El proceso está en memoria, pero no está usando CPU. Está esperando (sleeping) en la cola del scheduler."

**B. ¿Quién es su padre? (`parent`, `ppid`)**
En el PDF se menciona la jerarquía (Parent/Child).

```bash
cat status | grep PPid

```

* **Resultado:** Un número.
* **Actividad:** Haz `ps -p [NUMERO_DEL_PADRE]` y verán que es `bash`.
* **Lección:** "Bash creó a Python. Si mato a Bash, Python queda huérfano."

**C. La Memoria (`mm_struct`)**
En el PDF se habla de cómo el proceso gestiona su memoria.

```bash
cat maps

```

* **Resultado:** Un montón de direcciones hexadecimales.
* **Lección:** "Este es el mapa de memoria virtual del proceso. Aquí están el Heap, el Stack y las librerías cargadas."

---

### Parte 4: El Fenómeno "Zombie" (Muy Visual)

El PDF menciona la terminación de procesos y los estados. Vamos a crear un **Zombie** real. Esto siempre les gusta a los estudiantes.

1. Crea un archivo pequeño en C llamado `zombie.c`:
```c
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>

int main() {
    pid_t child_pid = fork();
    if (child_pid > 0) {
        // EL PADRE: Se duerme y NO recoge al hijo (no hace wait)
        printf("Soy el Padre (PID: %d). Mi hijo es %d.\n", getpid(), child_pid);
        printf("Mira 'ps' ahora. Mi hijo debería ser un ZOMBIE.\n");
        sleep(60); 
    } else {
        // EL HIJO: Termina inmediatamente
        printf("Soy el Hijo. Muero ya.\n");
        exit(0); 
    }
    return 0;
}

```


2. Compílalo y córrelo:
```bash
gcc zombie.c -o zombie
./zombie

```


3. Rápidamente, en la otra terminal, busca el proceso:
```bash
ps aux | grep Z

```


* **Resultado:** Verás algo como `[zombie] <defunct>` y en la columna de estado una **`Z`**.
* **Explicación Final:** "El hijo murió (`exit`), pero el padre sigue vivo y no ha leído su testamento (`wait`). El hijo ocupa espacio en la tabla de procesos (tiene PID) pero no consume memoria RAM. Es un muerto viviente en el Kernel."


