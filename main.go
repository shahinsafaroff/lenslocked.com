package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //response writer & set method
	fmt.Fprint(w,"<h1>Welcome to my awesome WebSite!</h1>")
}

func main () {
	http.HandleFunc("/",handlerFunc)
	http.ListenAndServe(":3000", nil)
}