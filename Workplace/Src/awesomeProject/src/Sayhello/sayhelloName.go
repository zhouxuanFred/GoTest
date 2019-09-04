package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("The r.form is  ", r.Body)
	fmt.Print("The Method is ",r.Host)
	fmt.Println("The body is", r.Header)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form{
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "hello Everyone!")
}

func start (){
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListernAndServer :", err)
	}
}

func main(){
	start()
}
