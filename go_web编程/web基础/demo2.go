package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func SayHi(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.Form)
	for k, v := range r.Form {
		fmt.Println(k)
		fmt.Println(v)
	}
	_, _ = fmt.Fprintf(w, "hello world")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request method", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(r.Form["username"])
		fmt.Println(r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", SayHi)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
