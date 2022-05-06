package main

import (
	"fmt"
    "bufio"
	"os"
	"strings"
	"strconv")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println(operation)
	operator := "p"
	values := strings.Split(operation, operator) //funcion para dividir string por una entrada
	fmt.Println(values)
	//fmt.Println(values[0] + values[1])
	operator1, err1 := strconv.Atoi(values[0]) //conversion con strconv, "_" la funcion retorna dos valores y esta manera de captura sin usuarlo
	fmt.Println("operator1:",operator1, "error",err1) //visualiar el error y la variable
	//uso del if para el control de errores
	if err1 != nil 	{
		fmt.Println("Error if",err1)
	} else {
		fmt.Println(operator1)
	}		//nil = null de retorno en la variable err1
	operator2, _ := strconv.Atoi(values[1])
	fmt.Println(operator1 + operator2) //se imprime la conversion
	switch operator {
	case "+":
		fmt.Println(operator1 + operator2)
	case "-":
		fmt.Println(operator1 - operator2)
	case "*":
		fmt.Println(operator1 * operator2)
	case "/":
		fmt.Println(operator1 / operator2)
	default: 
		fmt.Println("Don't understand this operator")
	}
}