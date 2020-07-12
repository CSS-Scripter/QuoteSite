package routes

import "github.com/gofiber/fiber"

// Read function reads the quotes from a file and returns the quotes
func Read(c *fiber.Ctx) {
	c.Send("Hello World")
}
