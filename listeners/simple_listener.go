package main

// TODO: Implement support for port rotation / listening to multiple ports and unifying data received across all of them

// TODO: Implement a relay listener which simply passes any data received onto another pre-configured network socket (IP or hostname and port)

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Based on code taken from https://astaxie.gitbooks.io/build-web-application-with-golang/en/03.2.html

func handleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself

	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

var listeningPort int

func init() {
	flag.IntVar(&listeningPort, "listeningPort", 9090, "Set the port to listen on; will use port 9090 by default")
}

func main() {
	fmt.Println("Listening for callback on port " + strconv.Itoa(listeningPort))

	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":" + strconv.Itoa(listeningPort), nil)
	
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
