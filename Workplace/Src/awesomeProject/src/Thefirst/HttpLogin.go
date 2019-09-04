package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func login (w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("helloword.html")
		t.Execute(os.Stdout,map[string]interface{}{"UserName": "helloword"})
	}else {
		r.ParseForm()
		s,_:=ioutil.ReadAll(r.Body)
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password", r.FormValue("password"))
		fmt.Println("body is :" , s)
		fmt.Println("header is : ", r.Header)
		fmt.Println("url is:", r.URL)
		fmt.Println("host is:", r.Host)
		fmt.Println("length", r.ContentLength)
	}
}
func login2 (w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		t,_ := template.ParseFiles("login.html")
		t.Execute(w,nil)
	}else {
		r.ParseForm()
		fmt.Println("username2:", r.Form["username"])
		fmt.Println("password2", r.Form["password"])
	}
}

func main(){
	http.HandleFunc("/login",login)
	http.HandleFunc("/login2",login2)
	err := http.ListenAndServe(":9091",nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
