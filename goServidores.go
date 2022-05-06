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

	for _, server := range servers{
		//trabajo asincrona
		go checkServer(server, canal) //se ejecuta de manera paralela o por hilos
	}

	//se utiliza este estilo de for pq el anterior debo usar todas las variables del for
	for i := 0; i < len(servers); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion %s\n", tiempoPaso)
}

func checkServer(server string, canal chan string){
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "Don't available")
		canal <- server + " no se encuentra disponible"
	}else{
		fmt.Println(server, "It's available")
		canal <- server + " se encuentra disponible"
	}
}