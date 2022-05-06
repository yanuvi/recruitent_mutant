package main

import "fmt"

type animal interface{
	move() string
}

type dog struct {}

type fish struct {}

type bird struct {}

func (dog) move() string {
	return "It's a dog and It's walking"
}

func (fish) move() string {
	return "It's a fish and It's swimming"
}

func (bird) move() string {
	return "It's a bird and It's flying"
}

func moveAnimal(a animal){
	fmt.Println(a.move())
}

func main () {
	d := dog{}
	moveAnimal(d)
	f := fish{}
	moveAnimal(f)
	b := bird{}
	moveAnimal(b)
}