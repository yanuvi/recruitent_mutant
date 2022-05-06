package main

import "fmt"

type task struct{
	name string
	description string
	complete bool
}

//*task es la variable de referencia a ajustar
func (t *task) markComplete () {
	t.complete = true
}

func (t *task) updateDescription (description string) {
	t.description = description
}

func (t *task) updateName (name string) {
	t.name =  name
}

func main(){
	t := task{
		name: "Completed my course of GO",
		description: "Completed my course of GO in this week", //siempre se debe colocar , al final de la variable
	}
	
	fmt.Printf("%+v\n", t) //Imprime una interface: propiedad + valor
	t.markComplete()
	t.updateName("Finish my course of GO")
	t.updateDescription("Finish my course of go inmediatly")
	fmt.Printf("%+v\n", t)
	
}
