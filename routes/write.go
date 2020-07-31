package routes

import (
	"encoding/json"
	"io/ioutil"

	"../structs"

	"github.com/gofiber/fiber"
)

// Write writes away an quote
func Write(c *fiber.Ctx) {
	c.Accepts("json", "text")
	c.Accepts("application/json")
	succes := writeToFile(c.Body())
	var message string
	if succes {
		message = `{"status":"200","message":"Quote created"}`
	} else {
		message = `{"status":"400","message":"Bad request"}`
	}
	c.Send(message)
}

func isQuoteValid(quote structs.Quote) bool {
	if quote.Message == "" || quote.By == "" || quote.Year == "" {
		return false
	}
	return true
}

func writeToFile(content string) bool {
	quotes := readQuotes()
	var quote structs.Quote
	err := json.Unmarshal([]byte(content), &quote)
	check(err)

	if !isQuoteValid(quote) {
		return false
	}

	quotes = append([]structs.Quote{quote}, quotes...)
	quoteString, err := json.Marshal(quotes)
	check(err)
	err = ioutil.WriteFile("quotes.json", []byte(quoteString), 0644)
	check(err)
	return true
}
