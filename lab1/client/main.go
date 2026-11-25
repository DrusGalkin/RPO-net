package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Клиент отправляет запрос...")

	resp, err := http.Get("http://127.0.0.1:9000/")
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Статус:", resp.Status)
	fmt.Println("Ответ:", string(body))
}
