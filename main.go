package main

import (
	"fmt"
	"net/http"

	"./config"
	"./controllers"
	_ "github.com/lib/pq"
)

var cfg = config.Config()

func main() {
	fmt.Println(cfg.Host)
	http.HandleFunc("/", controllers.SayHello)
	err := http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		fmt.Print("ListeAndServe: ", err)
	}
}
