package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main(){
	u,_ := url.Parse("http://localhost:9001/index")
	q := u.Query()
	q.Set("username","zhouxuan")
	q.Set("password","zhouxuan")

	u.RawQuery = q.Encode()
	res,err:= http.Get(u.String())
	fmt.Println(u.String())
	fmt.Println(res)
	if err != nil {
		log.Fatalln(err)
		return
	}

	result,err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("%s", result)
}