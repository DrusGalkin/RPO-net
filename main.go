package main

import (
	"fmt"
	"math/rand" // ДОБАВЛЕНО: для генерации случайного кода
	"strconv"   // ДОБАВЛЕНО: чтобы перевести число в строку

	"://github.com"
	"gopkg.in/mail.v2"
)

const PORT = 465
const SMTP_LINK = "smtp.mail.ru"
const MAIL_FROM = "твоя почта"
const MAIL_TO = "почта кому будем отправлять"
const PASSWORD = "пароль"

// ДОБАВЛЕНО: переменная для хранения сгенерированного кода в памяти сервера
var savedCode string

// ДОБАВЛЕНО: структура, чтобы прочитать JSON, который пришлет клиент в post-запросе
type ConfirmData struct {
	Code string `json:"code"`
}

func main() {
	app := fiber.New()

	// 1. GET-запрос: генерирует код и отправляет его на почту
	app.Get("/get_code", func(c fiber.Ctx) error { // ИЗМЕНЕНО: добавили имя 'c' и возврат 'error'
		// ДОБАВЛЕНО: генерируем случайное 4-значное число и сохраняем в переменную
		savedCode = strconv.Itoa(1000 + rand.Intn(9000))

		m := mail.NewMessage()
		m.SetHeader("From", MAIL_FROM)
		m.SetHeader("To", MAIL_TO)
		m.SetHeader("Subject", "Код подтверждения")
		m.SetBody("text/plain", "Ваш код: "+savedCode) // ИЗМЕНЕНО: подставили сгенерированный код

		d := mail.NewDialer(
			SMTP_LINK,
			PORT,
			MAIL_FROM,
			PASSWORD,
		)

		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Ошибка отправления!", err)
			return c.Status(500).SendString("Ошибка отправки письма") // ДОБАВЛЕНО: ответ клиенту при ошибке
		}

		fmt.Println("Письмо отправлено!")
		return c.SendString("Код отправлен на почту!") // ДОБАВЛЕНО: ответ клиенту при успехе
	})

	// 2. POST-запрос: принимает код от пользователя и сверяет его
	app.Post("/confirm", func(c fiber.Ctx) error { // ИЗМЕНЕНО: добавили 'c' и 'error'
		var data ConfirmData

		// ДОБАВЛЕНО: достаем JSON из тела запроса и кладем в структуру data
		if err := c.Bind().Body(&data); err != nil {
			return c.Status(400).SendString("Неверный формат данных")
		}

		// ДОБАВЛЕНО: сравниваем присланный код с тем, что мы сохранили при отправке
		if data.Code == savedCode && savedCode != "" {
			savedCode = "" // Сбрасываем код, чтобы его нельзя было ввести дважды
			return c.SendString("Код верный! Доступ разрешен.")
		}

		return c.Status(400).SendString("Неверный код!")
	})

	// ДОБАВЛЕНО (вместо цикла for): запускаем веб-сервер на порту 8080
	app.Listen(":8080")
}
