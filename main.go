package main

import (
	"fmt"
	"net/http"

	"./config"
	"./controllers"
	"./db"
)

var cfg = config.Config()

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Print("Database error", err)
	}
	http.HandleFunc("/", controllers.SayHello)
	fmt.Println("Server started on port", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), nil)
	if err != nil {
		fmt.Print("ListeAndServe: ", err)
	}

}
