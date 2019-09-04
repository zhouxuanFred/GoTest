package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

func login (w http.ResponseWriter, r *http.Request){
	fmt.Println("method", r.Method)
//	r.ParseForm()
	if r.Method == "GET" {

		//fmt.Println("method:", r.Method) //获取请求的方法
		//fmt.Println("http Header", r.Header)
		//result := loginValidation(r.FormValue("username"), r.FormValue("password"))
		//w.Write([]byte("true"))
	}else {
		var user map[string]interface{}
		s,_:=ioutil.ReadAll(r.Body)
		json.Unmarshal(s, &user)
		fmt.Println("获取json中的username:", user["username"].(string))
		fmt.Println("获取json中的password:", user["password"].(string))
		fmt.Println("",r.GetBody)
		fmt.Println("username:",r.PostFormValue("username"))
		fmt.Println("password:", r.PostFormValue("password"))
		fmt.Println("body is :" , s)
		//fmt.Println("header is : ", r.Header)
		//fmt.Println("url is:", r.URL)
		//fmt.Println("host is:", r.Host)
		//fmt.Println("length", r.ContentLength)
		LoginStatus:=loginValidation(user["username"].(string), user["password"].(string))
		if LoginStatus == 200{
			w.Header().Set("result","Username & password matched")
			w.WriteHeader(200)
		} else if (LoginStatus == 800){
			w.Header().Set("result","Username doesnt existed")
			w.WriteHeader(800)
		} else if (LoginStatus == 801){
			w.Header().Set("result","Password incorrect")
			w.WriteHeader(801)
		}
	}
}

func main(){
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9091",nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
