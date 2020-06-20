package main

import (
	"fmt"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Form)
}

func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe：", err)
	}

}
