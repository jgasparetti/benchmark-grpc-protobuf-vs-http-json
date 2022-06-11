package fiberjson

import (
	"errors"
	"log"
	"net/mail"

	"github.com/gofiber/fiber/v2"
)

// Start entrypoint
func Start() {
	app := fiber.New()
	app.Post("/", CreateUser)
	log.Println(app.Listen(":60002"))
}

// User type
type User struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Other    string    `json:"other"`
	Field1   int64     `json:"field1"`
	Field2   float64   `json:"field2"`
	Field3   []string  `json:"field3"`
	Field4   []int64   `json:"field4"`
	Field5   []float32 `json:"field5"`
}

// Response type
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	User    *User  `json:"user"`
}

// CreateUser handler
func CreateUser(c *fiber.Ctx) error {

	var user User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	validationErr := validate(user)
	if validationErr != nil {
		return c.Status(500).JSON(&Response{
			Code:    500,
			Message: validationErr.Error(),
		})
	}

	user.ID = "1000000"
	return c.Status(200).JSON(&Response{
		Code:    200,
		Message: "OK",
		User:    &user,
	})
}

func validate(in User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}
