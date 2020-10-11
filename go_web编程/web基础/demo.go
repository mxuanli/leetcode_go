package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("hello world")
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Println("-------")
	for k, v := range r.Form {
		fmt.Println("key：", k)
		fmt.Println("value：", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello astaxie!")
}

func main() {

	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe：", err)
	}

}
