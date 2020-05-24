package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

// Based on code taken from https://astaxie.gitbooks.io/build-web-application-with-golang/en/03.2.html

func handleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse arguments, you have to call this by yourself
	
	fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}


	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func main() {
    http.HandleFunc("/", handleRequest) // set router
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
