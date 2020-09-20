package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./config"
	_ "github.com/lib/pq"
)

var (
	cfg = config.Config()
	db  *sql.DB
	tml *template.Template
)

//Quote ...
type Quote struct {
	ID     int
	Author string
	Body   string
}

func main() {
	var err error

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgDB))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	tml = template.Must(template.ParseGlob("templates/*.gohtml"))

	http.HandleFunc("/", findAllQuote)
	http.HandleFunc("/add", addForm)
	http.HandleFunc("/add/process", addQuote)
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), nil)
}

func findAllQuote(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	rows, err := db.Query("SELECT * FROM quotes")
	if err != nil {
		http.Error(w, http.StatusText(405), http.StatusServiceUnavailable)
	}
	defer rows.Close()

	quotes := make([]*Quote, 0)
	for rows.Next() {
		quote := new(Quote)
		err := rows.Scan(&quote.ID, &quote.Author, &quote.Body)
		if err != nil {
			log.Fatal(err)
		}
		quotes = append(quotes, quote)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	tml.ExecuteTemplate(w, "quotes.gohtml", quotes)
}

func addForm(w http.ResponseWriter, req *http.Request) {
	tml.ExecuteTemplate(w, "add.gohtml", nil)
}

func addQuote(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	quote := Quote{}
	quote.Author = req.FormValue("author")
	quote.Body = req.FormValue("body")

	if quote.Author == "" || quote.Body == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO quotes (author, body) VALUES ($1, $2)", quote.Author, quote.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
