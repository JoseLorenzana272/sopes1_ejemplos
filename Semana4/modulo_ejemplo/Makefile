# Define el módulo de kernel objeto que se va a construir
obj-m += modulo_ejemplo.o

# Compila el módulo utilizando la infraestructura de compilación del kernel activo
all:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules

# Limpia los archivos temporales y binarios generados en la compilación
clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean