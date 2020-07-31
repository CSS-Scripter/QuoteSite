package structs

// Quote struct defines the structure of a quote
type Quote struct {
	Message string `json:"message"`
	By      string `json:"by"`
	Year    string `json:"year"`
}
