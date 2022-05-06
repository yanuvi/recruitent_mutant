package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

func CheckAuth() Middleware {
	return func (f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
			flag := true
			fmt.Println("Checking Autentication")
			if flag {
				f(w, r)
			}else {return}
		}
	}
}

func Logging() Middleware {
	return func (f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
			start := time.Now()
			defer func (){ //esta funcion anonima, sin nombre y no se repite
				log.Println(r.URL.Path, time.Since(start))	//Medir el tiempo de la ejecucion
			}()
			f(w,r)
		}
	}	
}