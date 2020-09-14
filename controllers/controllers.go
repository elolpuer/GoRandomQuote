package controllers

import (
	"fmt"
	"net/http"
)

//SayHello send hello
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
