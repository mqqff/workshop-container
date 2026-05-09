package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Workshop struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Speaker string `json:"speaker"`
	Date    string `json:"date"`
}

var db *sql.DB

func initDB() {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL not set")
	}

	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	fmt.Println("Connected to DB")
}

func getWorkshops(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT id, title, speaker, DATE_FORMAT(date, '%Y-%m-%d') FROM workshops`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var workshops []Workshop
	for rows.Next() {
		var wks Workshop
		err := rows.Scan(&wks.ID, &wks.Title, &wks.Speaker, &wks.Date)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		workshops = append(workshops, wks)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workshops)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func createWorkshop(w http.ResponseWriter, r *http.Request) {
	var wks Workshop

	err := json.NewDecoder(r.Body).Decode(&wks)
	if err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	result, err := db.Exec(
		"INSERT INTO workshops(title, speaker, date) VALUES(?, ?, ?)",
		wks.Title,
		wks.Speaker,
		wks.Date,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	wks.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(wks)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	initDB()

	http.HandleFunc("/workshops", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getWorkshops(w, r)
		case http.MethodPost:
			createWorkshop(w, r)
		default:
			http.Error(w, "Method not allowed", 405)
		}
	})

	fmt.Println("Server running on :8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
