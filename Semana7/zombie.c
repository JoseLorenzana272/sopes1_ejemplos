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