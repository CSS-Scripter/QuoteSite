package routes

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"../data"
	"../structs"

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

func readQuotes() []structs.Quote {
	path, err := os.Getwd()
	check(err)
	data, err := ioutil.ReadFile(path + "/quotes.json")
	check(err)
	quotesJSONString := string(data)
	var quotes []structs.Quote
	if err := json.Unmarshal([]byte(quotesJSONString), &quotes); err != nil {
		panic(err)
	}
	return quotes
}
