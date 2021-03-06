package controller

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/service"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	user, err := service.SignUp(data["username"], data["email"], data["password"], data["password_confirmation"])
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"error":   true,
			"code":    400,
			"message": err.Error(),
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token, err := service.SignIn(data["email"], data["password"])
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error":   true,
			"code":    400,
			"message": err.Error(),
		})
	}

	// create cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	}

	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
