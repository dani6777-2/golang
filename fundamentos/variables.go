package fundamentos

import "fmt"

/*
Variables en Go

Las variables son contenedores para almacenar valores de datos.

Tipos de datos básicos en Go:
  - bool: verdadero o falso (true/false)
  - numérico: enteros (int, int8, int16, int32, int64, uint, etc.)
             flotantes (float32, float64)
             complejos (complex64, complex128)
  - string: cadena de texto

De los tipos básicos se desprenden los demás tipos de datos compuestos:
  - arrays, slices, maps, structs, pointers, functions, interfaces, channels
*/

// Variables de nivel de paquete (package-level variables)
var (
	numero         int     = 3           // Almacena números enteros
	numeroFlotante float32 = 3.3         // Almacena números con punto flotante
	text           string  = "mi string" // Almacena texto
	boolean        bool    = true        // Almacena valores verdadero o falso (por defecto es false)
)

// PI es una constante: valor fijo e inmutable (solo lectura)
// Las constantes deben tener un valor asignado al momento de declararse
const PI = 3.1416

// DemoVariables demuestra el uso de diferentes tipos de variables y constantes
// También muestra las funciones de salida del paquete fmt
func DemoVariables() {
	fmt.Println("=== Demostración de Variables ===")
	fmt.Println()

	// Mostrar variables declaradas
	fmt.Println("Tipo de dato numérico int:", numero)
	fmt.Println("Tipo de dato número flotante:", numeroFlotante)
	fmt.Println("Tipo de dato cadena (string):", text)
	fmt.Println("Tipo de dato booleano:", boolean)
	fmt.Println("Tipo de variable CONSTANTE:", PI)
	fmt.Println()

	// Demostración de variables locales con inferencia de tipo
	fmt.Println("=== Variables Locales ===")
	nombre := "Juan"     // string inferido
	edad := 25           // int inferido
	altura := 1.75       // float64 inferido
	esEstudiante := true // bool inferido

	fmt.Printf("Nombre: %s, Edad: %d, Altura: %.2f, Es estudiante: %t\n",
		nombre, edad, altura, esEstudiante)
	fmt.Println()

	// Funciones de output en Go
	fmt.Println("=== Funciones de Salida en Go ===")
	fmt.Print("fmt.Print: imprime sin salto de línea")
	fmt.Print(" - continúa en la misma línea\n")
	fmt.Println("fmt.Println: imprime con salto de línea automático")

	// Verbos de formato comunes:
	// %v: valor por defecto, %T: tipo, %d: entero, %f: flotante, %s: string, %t: booleano
	fmt.Printf("fmt.Printf: formatea usando verbos como %%v (valor: %v) o %%T (tipo: %T)\n", numero, numero)
}
