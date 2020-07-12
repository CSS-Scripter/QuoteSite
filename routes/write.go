package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber"
)

// Write writes away an quote
func Write(c *fiber.Ctx) {
	c.Accepts("json", "text")
	c.Accepts("application/json")
	writeToFile(c.Body())
	c.Send(c.Body())
}

func writeToFile(content string) {
	quotes := readQuotes()
	var quote Quote
	err := json.Unmarshal([]byte(content), &quote)
	check(err)
	quotes = append([]Quote{quote}, quotes...)

	quoteString, err := json.Marshal(quotes)
	check(err)

	err = ioutil.WriteFile("quotes.json", []byte(quoteString), 0644)
	check(err)
}
