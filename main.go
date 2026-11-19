package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gofiber/fiber/v3"
	"gopkg.in/mail.v2"
)

const PORT = 465
const SMTP_LINK = "smtp.mail.ru"
const MAIL_FROM = "sergeym2900@mail.ru"
const MAIL_TO = "sergeym2900@gmail.com"
const PASSWORD = "GBzMeEZiDio6KiPVEhVq"
const MES = "Hello world!"

var (
	mu       sync.RWMutex
	lastCode string
)

func main() {
	app := fiber.New()

	app.Get("/get_code", func(c fiber.Ctx) error {
		if err := sendHello(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		mu.Lock()
		lastCode = MES
		mu.Unlock()

		return c.JSON(fiber.Map{
			"ok":      true,
			"message": MES,
		})
	})
	app.Post("/confirm", func(c fiber.Ctx) error {
		var req struct {
			Code string `json:"code"`
		}

		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "некорректный JSON",
			})
		}

		mu.RLock()
		expected := lastCode
		mu.RUnlock()

		valid := expected != "" && req.Code == expected

		return c.JSON(fiber.Map{
			"ok":    true,
			"valid": valid,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

func sendHello() error {
	m := mail.NewMessage()
	m.SetHeader("From", MAIL_FROM)
	m.SetHeader("To", MAIL_TO)
	m.SetHeader("Subject", "Hello")
	m.SetBody("text/plain", "Hello World!")

	d := mail.NewDialer(SMTP_LINK, PORT, MAIL_FROM, PASSWORD)
	d.SSL = true

	return d.DialAndSend(m)
}
