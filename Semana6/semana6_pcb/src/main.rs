// SIMULADOR DE PCB (Process Control Block) - SEMANA 6
// Objetivo: Entender cómo el SO estructura los procesos en memoria.

// 1. Enum para los Estados (Evita estados inválidos)
#[derive(Debug)]
enum EstadoProceso {
    Nuevo,
    Listo,
    Ejecutando,
    Terminado,
}

// 2. Struct que simula el 'task_struct' de Linux
#[derive(Debug)]
struct Proceso {
    id: u32,                // PID
    nombre: String,         // Nombre del proceso
    estado: EstadoProceso,  // Estado actual
    memoria_usada: u32,     // Stack + Heap simulado (KB)
}

impl Proceso {
    // Constructor (Como crear un proceso nuevo)
    fn new(id: u32, nombre: &str, memoria: u32) -> Proceso {
        Proceso {
            id,
            nombre: nombre.to_string(), // Copiamos el string al Heap
            estado: EstadoProceso::Nuevo,
            memoria_usada: memoria,
        }
    }

    // Simula al Scheduler dando CPU
    fn ejecutar(&mut self) {
        println!("--> [Scheduler] Cargando PID {}: {}", self.id, self.nombre);
        self.estado = EstadoProceso::Ejecutando;
    }
    
    // Simula la liberación de recursos (Exit)
    fn terminar(&mut self) {
        println!("--> [Kernel] Matando proceso: {}", self.nombre);
        self.estado = EstadoProceso::Terminado;
        self.memoria_usada = 0; // Memoria liberada
    }
}

fn main() {
    println!("=== INICIANDO KERNEL SIMULADO (RUST) ===");

    // 3. Vector que simula la Tabla de Procesos del Kernel
    let mut tabla_procesos: Vec<Proceso> = Vec::new();

    // Simulamos la creación de procesos (syscall fork)
    println!("\n[1] Creando procesos...");
    tabla_procesos.push(Proceso::new(1, "systemd", 2048));
    tabla_procesos.push(Proceso::new(102, "chrome_helper", 99999));
    tabla_procesos.push(Proceso::new(103, "vscode_server", 5000));

    // 4. Ciclo de Ejecución (Scheduler Round Robin simplificado)
    println!("\n[2] Iniciando Scheduler...");
    
    // Iteramos sobre referencias MUTABLES (&mut)
    for proceso in &mut tabla_procesos {
        
        // Paso 1: Poner en Listo
        proceso.estado = EstadoProceso::Listo;
        
        // Paso 2: Ejecutar
        proceso.ejecutar();
        println!("    * Estado: {:?} | Memoria: {} KB", proceso.estado, proceso.memoria_usada);
        
        // Simulación de tiempo de CPU...
        // (En un SO real aquí ocurriría la interrupción de reloj)
        
        // Paso 3: Terminar
        proceso.terminar();
        println!("    * Estado final: {:?}\n", proceso.estado);
    }
    
    println!("=== SISTEMA DETENIDO ===");
}