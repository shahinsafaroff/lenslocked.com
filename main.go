package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//func handlerFunc(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html") //response header writer & set method
//	//fmt.Fprint(w,r.URL.Path) //writes and reads the current path in web directory
//	if r.URL.Path =="/" { // this r means router
//		fmt.Fprint(w,"<h1>Welcome to my awesome WebSite!</h1>")
//	} else if r.URL.Path == "/contact" {
//		fmt.Fprint(w,"To get in touch, please send an email to a <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
//	} else {
//		w.WriteHeader(http.StatusNotFound) //Sets Status code 404 not found
//		fmt.Fprint(w, "<h1>We couldn't find the page you were looking for :-(</h1><p>Please email us if you keep being sent to an invalid page</p>")
//	}

//}

//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //For Third party library
//	fmt.Fprint(w, "hello, ", ps.ByName("name"))
//}
func home(w http.ResponseWriter, r *http.Request)    {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w,"<h1>Welcome to my awesome WebSite!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w,"To get in touch, please send an email to a <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Here's your Urgent Help Page\nPlease send an email if can't cope with FAQ<a href=\"mailto:faq@lenslocked.com\">faq@lenslocked.com</a>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,"<h1>Sorry But we couldn't the page you were looking for:(</h1>")
}

func main () {
	//mux := &http.ServeMux{}
	//mux.HandleFunc("/", handlerFunc) //adds more advanced router pattern / is default route to any subdirectory
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
	//router := httprouter.New() // Code from third party library
	//router.GET("/hello/:name/spanish", Hello) // Code from third party library
	//http.HandleFunc("/", handlerFunc)
	//http.ListenAndServe(":3000", nil)
	//Km@il Kraldi, Gazaxey de Gotdu!
}