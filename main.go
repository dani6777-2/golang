package main

import (
	"fmt"
	"learning-go/fundamentos"
	"learning-go/paquetes"
)

// main es la función principal donde inician las aplicaciones en Go
func main() {
	fmt.Println("═══════════════════════════════════════════")
	fmt.Println("    Aprendiendo Go - Fundamentos")
	fmt.Println("═══════════════════════════════════════════")
	fmt.Println()

	// Demostración de Variables
	fmt.Println("\n▶ VARIABLES Y TIPOS DE DATOS")
	fmt.Println("─────────────────────────────────────────")
	fundamentos.DemoVariables()

	// Demostración de Arrays
	fmt.Println("\n▶ ARRAYS (ARREGLOS)")
	fmt.Println("─────────────────────────────────────────")
	fmt.Println("\n1. Declaración de arrays:")
	fundamentos.Arrays()

	fmt.Println("\n2. Acceso a elementos:")
	fundamentos.AccessArrays()

	fmt.Println("\n3. Modificación de elementos:")
	fundamentos.ChangeArrays()

	fmt.Println("\n4. Longitud de arrays:")
	fundamentos.LengthArray()

	// Demostración de Slices
	fmt.Println("\n▶ SLICES (ARREGLOS DINÁMICOS)")
	fmt.Println("─────────────────────────────────────────")
	fmt.Println("\n1. Crear slice con literal:")
	fundamentos.CreateSliceLiteral()

	fmt.Println("\n2. Crear slice desde array:")
	fundamentos.CreateSliceFromArray()

	fmt.Println("\n3. Crear slice con make():")
	fundamentos.CreateSliceWithMake()

	fmt.Println("\n4. Operaciones con slices:")
	fundamentos.SliceOperations()

	// Demostración de Paquetes
	fmt.Println("\n▶ PAQUETES Y VISIBILIDAD")
	fmt.Println("─────────────────────────────────────────")
	paquetes.DemoPaquetes()

	fmt.Println("\n═══════════════════════════════════════════")
	fmt.Println("    Fin de la demostración")
	fmt.Println("═══════════════════════════════════════════")
}
