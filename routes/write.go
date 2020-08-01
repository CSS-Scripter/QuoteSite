package routes

import (
	"encoding/json"

	"../data"
	"../structs"

	"github.com/gofiber/fiber"
)

// WriteWithID writes to a quote to a server by it's server ID
func WriteWithID(c *fiber.Ctx) {
	c.Accepts("json", "text")
	c.Accepts("application/json")
	succes := createQuote(c.Params("id"), c.Body())
	var message string
	if succes {
		message = `{"status":"200","message":"Quote created"}`
	} else {
		message = `{"status":"400","message":"Bad request"}`
	}
	c.Send(message)
}

func createQuote(serverID string, content string) bool {
	var quote structs.Quote
	err := json.Unmarshal([]byte(content), &quote)
	check(err)

	if !isQuoteValid(quote) || len(serverID) >= 25 {
		return false
	}

	data.CreateQuote(serverID, quote)
	return true
}

func isQuoteValid(quote structs.Quote) bool {
	if quote.Message == "" || quote.By == "" || quote.Year == "" {
		return false
	}
	return true
}
