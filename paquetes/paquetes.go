package paquetes

/*
Paquetes en Go

Un paquete en Go es una colección de archivos de código fuente utilizados para
organizar y reutilizar código. Cada archivo de Go debe declarar un paquete al
principio del archivo.

Reglas de Visibilidad:
  - PÚBLICOS → Inician con letra MAYÚSCULA (se pueden usar desde otros paquetes)
  - privados → inician con letra minúscula (solo accesibles dentro del paquete)

Convenciones:
  - Todos los archivos en una misma carpeta deben pertenecer al mismo paquete
  - Los nombres de los paquetes deben ser en minúsculas, cortos y descriptivos
  - Evitar guiones bajos o nombres genéricos

Comandos importantes:
  - go mod init <nombre> : Inicializa un nuevo módulo de Go para gestionar dependencias
  - replace => Indica a Go que reemplace la dependencia por una carpeta local,
               en vez de buscar en el repositorio remoto

Alias de importación:
  - import io "fmt" : Se pueden dar alias a los paquetes para mayor claridad
                      o evitar conflictos de nombres
*/

import (
	io "fmt" // Ejemplo de alias: "io" en lugar de "fmt"
)

// secretFunction es una función privada (minúscula inicial)
// Solo puede ser llamada desde dentro del paquete "paquetes"
func secretFunction() {
	io.Println("  ❌ Función PRIVADA del paquete 'paquetes'")
	io.Println("     Solo accesible dentro de este paquete")
}

// PublicFunction es una función pública (mayúscula inicial)
// Puede ser llamada desde cualquier paquete que importe "paquetes"
func PublicFunction() {
	io.Println("  ✅ Función PÚBLICA del paquete 'paquetes'")
	io.Println("     Accesible desde otros paquetes")
}

// DemoPaquetes demuestra el uso de funciones públicas y privadas
// Esta es una función pública que orquesta la demostración
func DemoPaquetes() {
	io.Println("\nDemostración de visibilidad en paquetes:")
	io.Println()

	// Llamar a la función pública
	PublicFunction()
	io.Println()

	// Llamar a la función privada (solo funciona dentro del paquete)
	secretFunction()
	io.Println()

	io.Println("Nota: El alias 'io' se usa en lugar de 'fmt' en este paquete")
	io.Println("Esto demuestra cómo dar alias a las importaciones.")
}
