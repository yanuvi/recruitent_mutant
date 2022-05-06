package main

import "fmt"

func main(){
	m1 := make(map[string] int)

	m1["a"]  = 8

	fmt.Println(m1) 	 // imprime el valor de la llave
	fmt.Println(m1["a"]) // imprime el valor del mapa
}