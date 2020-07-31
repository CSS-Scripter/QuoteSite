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
	rows, err := db.Query("SELECT * FROM quotes WHERE server_id = $1", serverID)
	check(err)

	var results []structs.Quote

	defer rows.Close()
	for rows.Next() {
		var server string
		var quote string
		var by string
		var year string
		err = rows.Scan(&server, &quote, &by, &year)

		results = append(results, structs.Quote{
			Message: quote,
			By:      by,
			Year:    year,
		})
	}

	return results
}
