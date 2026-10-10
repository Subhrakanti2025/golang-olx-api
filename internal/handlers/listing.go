package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"uuid"
)

type listing struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type dbStruct struct {
	db *sql.DB
}

// that also the example of closure factory read about this as well closure
// func (d dbStruct) List(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

func NewDbConstructor(db *sql.DB) *dbStruct {
	return &dbStruct{db: db}
}
func (db dbStruct) List(w http.ResponseWriter, r *http.Request) {
	rows, err := db.db.Query(`SELECT id, title, description, price, city, created_at
			FROM listing ORDER BY created_at DESC
			LIMIT 200
			`)
	if err != nil {
		log.Printf("error from query %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	list := []listing{}
	for rows.Next() {
		var l listing
		err := rows.Scan(&l.Id, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt)
		if err != nil {
			log.Printf("error from process data %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		list = append(list, l)
	}
	if rows.Err() != nil {
		log.Printf("error from rows %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func (db dbStruct) DeleteList(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := db.db.Exec(`DELETE FROM listing where id=$1`, id)

	if err != nil {
		log.Printf("error from whene user delete some of information %v", err)
		http.Error(w, "inter server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
