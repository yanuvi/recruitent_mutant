package main

import (
	"fmt"
"net/http"
"time"
)

func main(){
	inicio := time.Now()
	servers := []string{
		"http://www.platzi.com",
		"http://www.google.com",
		"http://www.instagram.com",
		"http://www.facebook.com",
	}

	for _, server := range servers{
		checkServer(server)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion %s\n", tiempoPaso)
}

func checkServer(server string){
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "Don't available")
	}else{
		fmt.Println(server, "It's available")
	}
}