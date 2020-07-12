package routes

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber"
)

// Quote struct defines the structure of a quote
type Quote struct {
	Message string `json:"message"`
	By      string `json:"by"`
	Year    string `json:"year"`
}

// Index renders the index page
func Index(c *fiber.Ctx) {
	_ = c.Render("index", fiber.Map{
		"Quotes": readQuotes(),
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readQuotes() []Quote {
	path, err := os.Getwd()
	check(err)
	data, err := ioutil.ReadFile(path + "/quotes.json")
	check(err)
	quotesJSONString := string(data)
	var quotes []Quote
	if err := json.Unmarshal([]byte(quotesJSONString), &quotes); err != nil {
		panic(err)
	}
	return quotes
}
