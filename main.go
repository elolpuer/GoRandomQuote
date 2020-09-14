package main

import (
	"fmt"
	"net/http"

	"./controllers"
)

func main() {
	http.HandleFunc("/", controllers.SayHello)

	err := http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		fmt.Print("ListeAndServe: ", err)
	}
}
