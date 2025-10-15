package fundamentos

import "fmt"

//Variables en golang.
//son contenedores para almacenar valores de datos.

//Tipos de datos en go:

/*
basicos:
	- bool: verdadero o falso
	- numerico: enteros, flotantes y complejos
	-cadena: un valor de cadena

	**de los tipos basicos se desprenden los demas "sub-tipos" de datos.**

*/

var numero int = 3 //almacena numeros enteros.

var numeroFlotante float32 = 3.3 //almacena numeros con punto flotante.

var text string = "mi string" // almacena texto.

var boolean bool = true // almacena valores verdadero o falso, por defecto es false.

const PI = 3.1416 // valor fijo, inmutable solo lectura, se le debe asginar valor al declararse

func Helper() {

	//fmt: libreria de go para mostrar informacion por consola
	fmt.Println("Tipo de dato numerico int: ", numero)
	fmt.Println("Tipo de dato numero flotante: ", numeroFlotante)
	fmt.Println("Tipo de dato cadena:", text)
	fmt.Println("Tipo de dato booleano: ", boolean)
	fmt.Println("Tipo de variable CONSTANTE:", PI)

}
