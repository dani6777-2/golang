package fundamentos

import "fmt"

/*
Slices son estructuras de datos que permiten almacenar una colección de elementos
del mismo tipo. A diferencia de los arrays, los slices tienen un tamaño dinámico,
lo que significa que pueden crecer o reducirse según sea necesario.

Características principales:
  - Tamaño dinámico
  - Referencia a un array subyacente
  - Tienen longitud (len) y capacidad (cap)
*/

// Ejemplo de creación y uso de slices en Go

// CreateSliceLiteral demuestra la creación de un slice usando literales
// Sintaxis: slice_name := []dataType{values}
func CreateSliceLiteral() {
	mySlice := []int{1, 2, 3}

	fmt.Println("Slice creado con literal:", mySlice)
	fmt.Printf("Longitud: %d, Capacidad: %d\n", len(mySlice), cap(mySlice))
}

// CreateSliceFromArray demuestra cómo crear un slice a partir de un array
// Sintaxis: slice := array[inicio:fin]
func CreateSliceFromArray() {
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array original:", myArray)

	// Crear un slice que referencia los primeros 3 elementos del array
	mySlice := myArray[0:3]
	fmt.Println("Slice del array [0:3]:", mySlice)
	fmt.Printf("Longitud: %d, Capacidad: %d\n", len(mySlice), cap(mySlice))
}

// CreateSliceWithMake demuestra cómo crear un slice usando la función make()
// Sintaxis: slice_name := make([]type, length, capacity)
// - length: número de elementos inicializados a valor cero
// - capacity: capacidad total del slice (opcional)
func CreateSliceWithMake() {
	mySlice := make([]int, 4, 10)

	fmt.Println("Slice creado con make:", mySlice)
	fmt.Printf("Longitud: %d, Capacidad: %d\n", len(mySlice), cap(mySlice))
}

// SliceOperations demuestra las operaciones comunes con slices:
// - Acceso: acceder a elementos mediante índice
// - Modificación: cambiar el valor de un elemento
// - Append: agregar elementos al final
// - Concatenación: combinar múltiples slices
func SliceOperations() {
	pricesOld := []int{300, 500, 6000, 7000}

	// Acceso
	fmt.Println("Acceso al elemento en índice 2:", pricesOld[2])

	// Modificar
	pricesOld[2] = 5000
	fmt.Println("Después de modificar índice 2:", pricesOld)

	// Agregar elementos usando append
	pricesOld = append(pricesOld, 8000, 90000)
	fmt.Println("Después de agregar elementos:", pricesOld)

	// Concatenar slices usando el operador ... (spread)
	pricesNew := []int{10000, 40000, 50000, 20000, 50000}
	allPrices := append(pricesOld, pricesNew...)
	fmt.Println("Slices concatenados:", allPrices)
	fmt.Printf("Longitud final: %d, Capacidad: %d\n", len(allPrices), cap(allPrices))
}
