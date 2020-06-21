package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //response header writer & set method
	//fmt.Fprint(w,r.URL.Path) //writes and reads the current path in web directory
	if r.URL.Path =="/" { // this r means router
		fmt.Fprint(w,"<h1>Welcome to my awesome WebSite!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w,"To get in touch, please send an email to a <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	}

}

func main () {
	http.HandleFunc("/",handlerFunc)
	http.ListenAndServe(":3000", nil)
}