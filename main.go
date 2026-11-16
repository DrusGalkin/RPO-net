package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"gopkg.in/mail.v2"
)

const PORT = 465
const SMTP_LINK = "smtp.mail.ru"
const MAIL_FROM = "твоя почта"
const MAIL_TO = "почта кому будем отправлять"
const PASSWORD = "пароль"

func main() {
	app := fiber.New()

	app.Get("/get_code", func(fiber.Ctx) {

	})
	app.Post("/confirm", func(fiber.Ctx) {

	})

	for {
		m := mail.NewMessage()
		m.SetHeader("From", MAIL_FROM)
		m.SetHeader("To", MAIL_TO)
		m.SetHeader("Subject", "HTML-письмо с вложением")

		m.SetBody("text/plain", "Если вы видите этот текст, ваш клиент не поддерживает HTML.")
		m.AddAlternative("text/html", "<h1>Привет!</h1><p>Это HTML-версия письма.</p>")

		d := mail.NewDialer(
			SMTP_LINK,
			PORT,
			MAIL_FROM,
			PASSWORD,
		)
		d.StartTLSPolicy = mail.MandatoryStartTLS

		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Ошибка отпавления!", err)
		} else {
			fmt.Println("Письмо отправлено!")
		}

		time.Sleep(60 * time.Second)
	}

}
