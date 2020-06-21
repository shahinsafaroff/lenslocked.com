package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //response header writer & set method
	//fmt.Fprint(w,r.URL.Path) //writes and reads the current path in web directory
	if r.URL.Path =="/" { // this r means router
		fmt.Fprint(w,"<h1>Welcome to my awesome WebSite!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w,"To get in touch, please send an email to a <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound) //Sets Status code 404 not found 
		fmt.Fprint(w, "<h1>We couldn't find the page you were looking for :-(</h1><p>Please email us if you keep being sent to an invalid page</p>")
	}

}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //For Third party library
	fmt.Fprint(w, "hello, ", ps.ByName("name"))
}

func main () {
	//mux := &http.ServeMux{}
	//mux.HandleFunc("/", handlerFunc) //adds more advanced router pattern / is default route to any subdirectory
	router := httprouter.New() // Code from third party library
	router.GET("/hello/:name/spanish", Hello) // Code from third party library
	//http.HandleFunc("/",handlerFunc)
	http.ListenAndServe(":3000", router)
}