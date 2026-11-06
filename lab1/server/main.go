package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--------------- Sever started ---------------")

	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	s := bufio.NewScanner(c)
	for s.Scan() {
		if _, err := fmt.Fprintln(c, "Запрос на сервер: "+c.LocalAddr().String()+" отправлен"); err != nil {
			log.Print(err)
			return
		}
		fmt.Println("Новое запрос с адреса - " + s.Text())
	}

	if err := s.Err(); err != nil {
		log.Print(err)
	}
}
