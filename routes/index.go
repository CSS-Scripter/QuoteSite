package routes

import (
	"../data"

	"github.com/gofiber/fiber"
)

// Index renders the index page
func Index(c *fiber.Ctx) {
	_ = c.Render("index", fiber.Map{})
}

// IndexWithID renders the quotes from a certain server
func IndexWithID(c *fiber.Ctx) {
	_ = c.Render("index", fiber.Map{
		"Quotes": data.GetQuotesFromServer(c.Params("id")),
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
