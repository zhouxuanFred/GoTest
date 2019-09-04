package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func loginValidation (username,password string) int{
	fmt.Println("This is a program to validate login_info: %s, %s", username,password)

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connection isses", err.Error())
	}

	defer c.Close()
//	_, err = c.Do("SET", "mykey", "superWang")
//	if err != nil {
//		fmt.Println("redis set failed:", err)
//	}
	fmt.Println("validate username and password")
	password_redis, err := redis.String(c.Do("GET",username))
	fmt.Println("password :  %s",password_redis)
	if err != nil {
		// no registion data
		fmt.Println("redis get failed:", err)
		return 800   // no username
	} else {
		fmt.Printf("Get password: %v \n", password_redis)
	}

	if (password == password_redis) {
		return 200   // pass
	} else {
		return 801   // incorrect password
	}
}

func Updatepassword(password string)(int){
	fmt.Print("This is the password update action")

	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connection issues", err.Error())
	}
	defer c.Close()

	return 0
}