package main

import (
	"fmt"
"net/http"
"time"
)

func main(){
	inicio := time.Now()
	canal := make(chan string)
	servers := []string{
		"http://www.platzi.com",
		"http://www.google.com",
		"http://www.instagram.com",
		"http://www.facebook.com",
	}
	i := 0
	for {
		if i > 2 {
			break
		}
		for _, server := range servers{
			//trabajo asincrona
			go checkServer(server, canal) //se ejecuta de manera paralela o por hilos
		}
		time.Sleep(1*time.Second)
		fmt.Println(<-canal)
		i++
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion %s\n", tiempoPaso)
}

func checkServer(server string, canal chan string){
	_, err := http.Get(server)
	if err != nil {		
		canal <- server + " no se encuentra disponible"
	}else{		
		canal <- server + " se encuentra disponible"
	}
}