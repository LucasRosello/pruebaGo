package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundoooooo2")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint de la API")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}

	fmt.Println(metadata["name"])
	fmt.Fprintf(w, "Payload %v\n", metadata)
}
