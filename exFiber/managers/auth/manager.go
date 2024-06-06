package auth

import (
	"fmt"
	"time"

	"middleware/session"

	"github.com/gofiber/fiber/v2"
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

		loginTime := ssn.Get("LoginID")
		if loginTime == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		fmt.Printf("[DEBUG] LoginID   : %v\n", ssn.Get("LoginID"))
		fmt.Printf("[DEBUG] Password  : %v\n", ssn.Get("Password"))
		fmt.Printf("[DEBUG] Login-Time: %v\n", ssn.Get("Login-Time"))
		fmt.Printf("[DEBUG] AID       : %v\n", ssn.Get("AID"))

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

	ssn.Set("LoginID", req.LoginID)
	ssn.Set("Password", req.Password)
	ssn.Set("Login-Time", time.Now().Unix())
	ssn.Set("AID", 777)

	if err := ssn.Save(); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
