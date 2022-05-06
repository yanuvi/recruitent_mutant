package main

import "fmt"

func main(){
	x := 25
	fmt.Println(x)
	fmt.Println(&x) //valor de la instancia o referencia en memoria
	y := &x	//asignacion de la instancia o referencia
	fmt.Println(y)
	fmt.Println(*y) //valor de la variable
	changeValue(x)
}

func changeValue(a int){
	a = 36
}