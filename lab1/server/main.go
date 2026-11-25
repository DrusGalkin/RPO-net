package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

var whiteList = map[string]bool{
	"127.0.0.1": true,
}

func isIPAllowed(ip string) bool {
	return whiteList[ip]
}

func handler(w http.ResponseWriter, r *http.Request) {
	clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Ошибка получения IP", http.StatusInternalServerError)
		return
	}

	if !isIPAllowed(clientIP) {
		log.Println("Отказано в доступе для IP:", clientIP)
		http.Error(w, "403: Доступ запрещен", http.StatusForbidden)
		return
	}

	log.Println("Успешное подключение от IP:", clientIP)
	fmt.Fprintf(w, "Привет! Твой IP (%s) в белом списке.", clientIP)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущен на порту 9000")
	if err := http.ListenAndServe("127.0.0.1:9000", nil); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}
}
