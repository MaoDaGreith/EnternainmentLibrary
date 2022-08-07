package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/contact" {
		fmt.Fprint(w, "For more info, please contact me at <a "+
			"href=\"mailto:traianlupanciuc@gmail.com\">traianlupanciuc@gmail.com</a>")
	} else if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>This is a title Page</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>The Page you are looking for does not exist :(</h1>")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/contact", handlerFunc)
	http.ListenAndServe(":3000", r)
}
