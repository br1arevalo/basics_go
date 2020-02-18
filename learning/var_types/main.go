package main

import "fmt"

func main() {

	//Variables
	//A var declaration creates a variable of a particular type, attaches a name to it, and sets its initial value
	// var nombreVar tipoVar = valor
	// nombreVar := valor
	// nombreVar1, nombreVar2 := valor1, valor 2

	//Tipos de variables

	//string
	var strings string = "String es un tipo de variable que sirve para guardar una cadena de caracteres"
	fmt.Println(strings)

	cadena := "sirve declarado de diversas formas como: nombreVar := valor"
	fmt.Println(cadena)

	str1, str2 := "Se pueden declarar varias variables en una misma ", "declaracion "
	fmt.Println(str1, str2)

	//bool
	st, boo := "Bool se declara con el tipo de variable, o con el valor inicial que se indique, por ejemplo: ", true
	fmt.Println(st, boo)

	//se puede modificar
	boo = false
	fmt.Println(boo)

	// int
	s, in := "Int es para guardar datos numericos enteros como", 69
	fmt.Println(s, in)

	//Quiero terminar todos los tipos de variables

}
