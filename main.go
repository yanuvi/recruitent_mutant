package main //SIEMPRE VA EL NOMBRE DEL PAQUETE

import "fmt" //MANERA DE IMPORTAR SOLO UNA LIBRERIA

//func main(){
	//fmt.Println("Hello World Platzi")
//}

func main(){
	 var mensaje string = "Hello World"
	 mensajeFacil := "Hello World used :=" //esta manera el lenguaje infiere la declaracion
	 fmt.Println(mensaje)
	 fmt.Println((mensajeFacil))
	 a:= 10.
	 var b float64 = 3
	 fmt.Println((a / b))
	 var c int = 10
	 d := 3
	 fmt.Println(c /d)
	 x := true
	 y := false
	 fmt.Println((x || y)) // operador or
	 fmt.Println(y && x) //operador and
	 fmt.Println(!x) //Manera de negar una variable booleana
	 privada() //solo lo puede utilizar el paquete que la contenga
	 Publica() //Debe estar en mayuscula la primera letra, publica o exportada
	 imprimirHola()
}

func privada() {
	fmt.Println("Logic without export")
}

func Publica() {
	fmt.Println("Logic with export")
}

func imprimirHola() {
	defer fmt.Println("World") //Esto se hara antes de terminar la funcion
	fmt.Println("Hello")
}