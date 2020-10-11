package main

import (
	"fmt"
	"log"
	"net/http"
)

func SayHi(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form)
	//fmt.Println(r.)
	_, _ = fmt.Fprintf(w, "<h1>hello world</h1>")

}

func main() {
	http.HandleFunc("/", SayHi)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
