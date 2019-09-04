package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is an example of https service in Goland")
}


func handlerClientConnnection(){

}

func main()  {
	// http.HandleFunc("/",handler)





	err := http.ListenAndServeTLS(":9091","/Users/zhouxuan/Documents/CATest/server.crt","/Users/zhouxuan/Documents/CATest/server.key",nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}