package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("HTMLs/View/index.html")
	t.Execute(w, nil)
}


func main() {
	http.HandleFunc("/", Index)
	// 监听本机的8080端口
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
