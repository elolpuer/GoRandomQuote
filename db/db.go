package db

import (
	"database/sql"
	"fmt"
	"log"

	"../config"
	"../models"
	_ "github.com/lib/pq"
)

var cfg = config.Config()

var db sql.DB

//Connect ...
func Connect() error {
	var err error

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgDB))
	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

// FindAllQuote ...
func FindAllQuote() {
	rows, err := db.Query(`SELECT * FROM quotes`)
	if err != nil {
		fmt.Println("Query to quotes", err)
	}
	defer rows.Close()

	quotes := make([]*models.Quote, 0)
	for rows.Next() {
		quote := new(models.Quote)
		err := rows.Scan(&quote.ID, &quote.Author, &quote.Body)
		if err != nil {
			log.Fatal(err)
		}
		quotes = append(quotes, quote)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, quote := range quotes {
		fmt.Printf("%d %s %s\n", quote.ID, quote.Author, quote.Body)
	}

}
