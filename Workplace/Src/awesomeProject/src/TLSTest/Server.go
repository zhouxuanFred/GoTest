package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func CatchErrorlog(){

}

func HandeClientConnnet(conn net.Conn){
	defer conn.Close()
	fmt.Println("Recieved Connection from", conn.RemoteAddr().String())

	buffer := make([]byte, 1024)

	for{
		len, err := conn.Read(buffer)
		if err != nil {
			log.Fatalln(err.Error())
			break
		}

		fmt.Println("Recieved data: %s\n", string(buffer[:len]))
		_, err = conn.Write([]byte("Server recived DATA:" + string(buffer[:len])))

		if err != nil {
			break
		}
	}

	fmt.Println("Client" + conn.RemoteAddr().String() + "Connection closed..")
}


func main(){
	crt, err := tls.LoadX509KeyPair("/Users/zhouxuan/Documents/CATest/server.crt", "/Users/zhouxuan/Documents/CATest/server.key")
	if err != nil {
		log.Fatalln(err.Error())
	}

	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	tlsConfig.InsecureSkipVerify = true

	fmt.Println(*tlsConfig)

	l, err := tls.Listen("tcp", ":8888", tlsConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			go HandeClientConnnet(conn)
		}
	}
}

