package main

import "fmt"

func main() {

	//Declarations
	//A declaration names a program entity and specifies some or all of its properties
	//There are four major kinds of declarations: var, const, type, and func

	//Ejemplos
	//A las variables se les puede actualizar su valor en cualquier momento

	//Declaracion general funcional en cualquier lado del codigo
	var br1 string = "Se puede declarar afuera de una funcion y tambien dentro con la declaracion var nombreVar tipoVar = valor"

	fmt.Println(br1)

	br1 = "Podemos actualizar la variable con nombreVar = nuevoValor"

	fmt.Println(br1)

	//Declaracion rapida solo funcional dentro de alguna funcion
	br2 := "Solo se puede declarar dentro de una funcion, estamos dentro de la funcion ejemplo, por lo tanto sirve la declaracion br2 := valor"

	fmt.Println(br2)

	///A las constantes no se les puede modificar

	const constante = "Solo podemos declarar una vez una constante con const nombreConst = valor"

	fmt.Println(constante)

	declaration()

}

//A function declaration has a name, a list of parameters (the variables whose values are
//provided by the function’s callers), an optional list of results, and the function body, which
//contains the statements that define what the function does

func declaration() {

	//Esta parte del codigo va a ser actualizada, tenia mejor codigo, solo que me atore en algo y lo hice mas literal para que se viera como queria
	//igual para darme cuenta si leiste todo, si estas leyendo esto mandame alguna frase del payaso eso por whats
	//si no, te cagaste mijo, esto ya quedo en la historia de mi github

	cuerpoFuncion := "Las funciones se declaran de la siguiente manera:"
	fmt.Println(cuerpoFuncion)
	e1 := "func nombreFuncion (lista_de_parametros tipo_de_parametros) lista_de_resultados { "
	fmt.Println(e1)
	e2 := "cuerpo_de_la_funcion "
	fmt.Println(e2)
	e3 := "}"
	fmt.Println(e3)

}

//func "nombreFuncion" ("lista de parametros" "tipo de parametros") "lista de resultados" {
//	"cuerpo de la funcion"
//}
