package controllers

import (
	"fmt"
	"net/http"
	// "../db"
)

//SayHello send hello
func SayHello(w http.ResponseWriter, r *http.Request) {
	// db.FindAllQuote()
	fmt.Fprint(w, "Hello")
}
