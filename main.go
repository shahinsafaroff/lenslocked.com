package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visited our page")
	fmt.Fprint(w,"<h1>Welcome to my awesome WebSite!</h1>")
}

func main () {
	http.HandleFunc("/",handlerFunc)
	http.ListenAndServe(":3000", nil)
}