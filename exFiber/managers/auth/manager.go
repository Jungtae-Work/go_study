package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
)

type loginRequest struct {
	LoginID  string `form:"loginid"`
	Password string `form:"password"`
}

var store = session.New(session.Config{
	Storage:   redis.New(),
	KeyLookup: "header:Session-ID",
})

func Set(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if c.OriginalURL() == "/login" {
			return c.Next()
		}

		ssn, err := store.Get(c)
		if err != nil {
			return err
		}

		loginTime := ssn.Get("login")
		if loginTime == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("AID", ssn.Get("AID"))

		return c.Next()
	})
}

func Login(c *fiber.Ctx) error {
	req := loginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if req.LoginID != "viper" || req.Password != "12345" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	ssn, err := store.Get(c)
	if err != nil {
		return err
	}

	ssn.Set("login", time.Now().Unix())
	ssn.Set("AID", 777)

	if err := ssn.Save(); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
