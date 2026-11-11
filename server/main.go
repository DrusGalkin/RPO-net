package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// список допустимых адресов
	allowedIp := []string{"127.0.0.1"}

	// прослушивание айпи и порта(топтать его в рот )
	ln, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--------------- Server started ---------------")

	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// в переменную host записываем чисто вылизанный айпишник входящего запроса
		host, _, err := net.SplitHostPort(c.RemoteAddr().String())

		//если айпишник не в белом списке,закрываем соеденение
		if !isAllowed(host, allowedIp) {
			fmt.Println("Доступ запрещен для IP:", c.RemoteAddr().String())
			c.Close()
			continue
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

// функция проверки разрешенного ip
func isAllowed(host string, allowedIp []string) bool {
	for _, ip := range allowedIp {
		if ip == host {
			return true

		}
	}
	return false

}
