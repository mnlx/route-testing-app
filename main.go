package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":9999", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path[1:])
	fmt.Fprintf(w, "This is version 2. Hello, %s!", r.URL.Path[1:])
}