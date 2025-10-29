package fundamentos

import "fmt"

/*
Arrays (Arreglos) en Go son colecciones de elementos del mismo tipo con un tamaño fijo.
Una vez que se define el tamaño de un arreglo, no se puede cambiar.

La sintaxis para declarar un arreglo es la siguiente:

	var nombreArreglo [tamaño]tipoDeDato

Ejemplo:
	var numeros [5]int -> declara un arreglo de enteros con capacidad para 5 elementos

Los arreglos en Go son valores, lo que significa que cuando se asigna un arreglo a otra variable
o se pasa a una función, se crea una copia del arreglo original.
*/

// Arrays demuestra la declaración básica de arreglos en Go
func Arrays() {
	var arrayUno = [4]int{1, 2, 3, 4}
	arrayDos := [4]int{5, 6, 7, 8}

	fmt.Println("Array uno:", arrayUno)
	fmt.Println("Array dos:", arrayDos)
}

// AccessArrays demuestra cómo acceder a elementos individuales de un arreglo
func AccessArrays() {
	var arrayUno = [4]int{1, 2, 3, 4}
	arrayDos := [4]int{5, 6, 7, 8}

	fmt.Println("Elemento en índice 1 del array uno:", arrayUno[1])
	fmt.Println("Elemento en índice 0 del array dos:", arrayDos[0])
}

// ChangeArrays demuestra cómo modificar elementos de un arreglo
func ChangeArrays() {
	var arrayUno = [4]int{1, 2, 3, 4}
	arrayDos := [4]int{5, 6, 7, 8}

	fmt.Println("Arrays antes de modificar:")
	fmt.Println("Array uno:", arrayUno)
	fmt.Println("Array dos:", arrayDos)

	arrayUno[1] = 30
	arrayDos[3] = 40

	fmt.Println("\nArrays después de modificar:")
	fmt.Println("Array uno:", arrayUno)
	fmt.Println("Array dos:", arrayDos)
}

// LengthArray demuestra cómo obtener la longitud de un arreglo usando len()
func LengthArray() {
	var arrayUno = [4]int{1, 2, 3, 4}
	arrayDos := [4]int{5, 6, 7, 8}

	fmt.Println("Longitud del array uno:", len(arrayUno))
	fmt.Println("Longitud del array dos:", len(arrayDos))
}
