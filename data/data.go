package data

import (
	"database/sql"
	"fmt"

	"../structs"

	// Blanc import because
	_ "github.com/lib/pq"
	"github.com/tkanos/gonfig"
)

// Configuration structure defines the information written in the config file
type Configuration struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var configuration Configuration
var psqlInfo string

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// InitConfig initializes the configuration for the databaseconnection
func InitConfig() {
	configuration = Configuration{}
	err := gonfig.GetConf("./config.json", &configuration)
	check(err)

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configuration.Host, configuration.Port, configuration.User, configuration.Password, configuration.Database)
}

func openConnection() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	check(err)
	err = db.Ping()
	check(err)
	return db
}

// GetQuotesFromServer takes a serverID, and returns all quotes from the given server
func GetQuotesFromServer(serverID string) []structs.Quote {
	db := openConnection()
	defer db.Close()
	rows, err := db.Query("SELECT quote, by, year FROM quotes WHERE server_id = $1 ORDER BY created_at DESC", serverID)
	check(err)

	var results []structs.Quote

	defer rows.Close()
	for rows.Next() {
		var quote, by, year string
		err = rows.Scan(&quote, &by, &year)

		results = append(results, structs.Quote{
			Message: quote,
			By:      by,
			Year:    year,
		})
	}

	return results
}

// CreateQuote creates an quote for a certain server
func CreateQuote(serverID string, quote structs.Quote) {
	db := openConnection()
	defer db.Close()
	insertStatement := `INSERT INTO quotes(server_id, quote, by, year) VALUES($1, $2, $3, $4)`
	db.QueryRow(insertStatement, serverID, quote.Message, quote.By, quote.Year)
}
