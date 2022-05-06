package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func HandleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World!")
}

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is the API Endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w,"Error: %v", err)
		return
	}
	fmt.Fprintf(w,"Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w,"Error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}