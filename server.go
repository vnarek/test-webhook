package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int = 0

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	count++
	log.Println("got requeste", r, " number", count)
	if count > 5 {
		go panic("can't handle this many requests!")
	}
	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		log.Println(err.Error())
	}
}
