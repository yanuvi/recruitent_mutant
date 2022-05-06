package main

import "fmt"

type taskList struct{
	tasks []*task//arreglo dinamico o slices, apuntando con la task ya definido
}

func (t *taskList) addList (tl *task) { //* es para el uso de las referencias de las variables 
	t.tasks = append (t.tasks, tl) //1er parametro es la lista, 2do lo q se va a agregar
}

func (t *taskList) deleteList (index int){
	// Se crea un elemento nuevo solo con los elementos de la izquierda del indice
	// y todos los elementos de la derecha del indice
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...) 
	//t.tasks[:index] - todos los elementos de la izquierda 
	//t.tasks[index+1:]... - todos los elmentos de la derecha seguido del indice (+1)
}

func (t *taskList) imprimeList(){	
	for _, task := range t.tasks {
		fmt.Println("Name", task.name)
		fmt.Println("Description", task.description)
	}
	  
}

func (t *taskList) imprimeListCompleted(){	
	for _, task := range t.tasks {
		if task.complete {
			fmt.Println("Name", task.name)
			fmt.Println("Description", task.description)
		}
	}
}

func (t *task) markComplete () {
	t.complete = true
}

type task struct{
	name string
	description string
	complete bool
}

func main(){
	//& valor de la referencia o instancia de la variable
	t1 := &task{
		name: "Completed my course of GO",
		description: "Completed my course of GO in this week", //siempre se debe colocar , al final de la variable
	}
	t2 := &task{
		name: "Completed my course of Python",
		description: "Completed my course of Python in this week", //siempre se debe colocar , al final de la variable
	}
	t3 := &task{
		name: "Completed my course of NodeJS",
		description: "Completed my course of NodeJS in this week", //siempre se debe colocar , al final de la variable
	}

	//metodo para usar las tareas
	list := &taskList {
		tasks: []*task{ //Se agregan los elementos de la tarea
			t1, t2,
		},
	}

	fmt.Println(list.tasks[0])
	list.addList(t3)
	fmt.Println(len(list.tasks)) //funcion para medir la longitud del slice
	//list.deleteList(1) //recibe el valor de un indice o posicion del slice
	//fmt.Println(len(list.tasks))
	fmt.Println("Func imprimeList")
	
	for i:= 0; i < len(list.tasks); i++ {		
		fmt.Println("for ve - Index", i, "nombre", list.tasks[i].name)
	}

	// esta es la manera mas utilizada del for
	// donde se realiza el for por el tamano o rango del slice 
	// sin declarar otras varialbes
	for index, task := range list.tasks {
		fmt.Println("For range - Index", index, "nombre", task.name)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
			//utilizado para romper el ciclo al cumplir la condicion
		}
		//fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
			//utilizado para omitir esta condicion y continua
		}
		//fmt.Println(i)
	}

	list.imprimeList()
	list.tasks[0].markComplete()
	fmt.Println("Task completed")
	list.imprimeListCompleted()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Yanuvi"] = list

	t4 := &task{
		name: "Completed my course of Java",
		description: "Completed my course of Java in this week", //siempre se debe colocar , al final de la variable
	}
	t5 := &task{
		name: "Completed my course of C#",
		description: "Completed my course of C# in this week", //siempre se debe colocar , al final de la variable
	}

	list2 := &taskList {
		tasks: []*task{ 
			t4, t5,
		},
	}

	mapaTareas["Yenny"] = list2

	fmt.Println("Tareas de Yanuvi")
	mapaTareas["Yanuvi"].imprimeList()
	fmt.Println("Tareas de Yenny")
	mapaTareas["Yenny"].imprimeList()
}
