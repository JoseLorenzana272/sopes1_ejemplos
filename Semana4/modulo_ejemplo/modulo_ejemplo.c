// headers para la cración del modulo de kernel
#include <linux/module.h> // funciones y macros para módulos
#include <linux/kernel.h> // funciones para imprimir mensajes en el kernel
#include <linux/init.h>   // macros para las funciones de inicialización y limpieza

// Info del modulo
MODULE_LICENSE("GPL"); // Tipo de licencia (obligatorio)
MODULE_AUTHOR("Jose"); // Autor del módulo
MODULE_DESCRIPTION("Un modulo de ejemplo para el curso de sistemas operativos"); 

// función que se ejcuta cuando carguemos el modulo
static int __init modulo_inicio(void)
{
    printk(KERN_INFO "Modulo cargado Yupiii\n"); // imprimir mensaje en el log del sistema
    printk(KERN_INFO "HOLAAAA; El modulo se cargo bien\n");
    return 0; // Indica que la carga fue exitosa
}

// se ejecuta cuando el modulo se elimina
static void __exit modulo_fin(void)
{
    printk(KERN_INFO "Modulo eliminado\n"); // imprimir mensaje en el log del sistema
    printk(KERN_INFO "ADIOS, el modulo se eliminó correctamente\n");
}

// macros que indican cual es la funcion de inicio y salida
module_init(modulo_inicio); // Al cargar el módulo, se llama a modulo_inicio
module_exit(modulo_fin); // Al eliminar el módulo, se llama a modulo_fin

