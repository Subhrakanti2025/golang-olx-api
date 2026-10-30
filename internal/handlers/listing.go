package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"time"
	"uuid"

	"github.com/Subhrakanti2025/golang-olx-api/internal/httpx"
	"github.com/Subhrakanti2025/golang-olx-api/internal/middleware"
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
	db     *sql.DB
	logger *slog.Logger
}

// that also the example of closure factory read about this as well closure
// func (d dbStruct) List(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

func NewDbConstructor(db *sql.DB, logger *slog.Logger) *dbStruct {
	return &dbStruct{
		db:     db,
		logger: logger,
	}
}
func (db dbStruct) List(w http.ResponseWriter, r *http.Request) {
	// request scoped context
	// Get the context of the incoming HTTP request.
	// If the client cancels the request (e.g., closes the page,
	// reloads it, or stops	 the request in Postman), the context
	// is canceled when Go detects the disconnection.
	// Database queries using this context can then stop,
	// avoiding unnecessary work.
	ctx := r.Context()

	// GetRequestId retrieves the request ID stored in the context by our middleware.
	//
	// We use a custom, unexported context key type (`ctxType`) and an
	// unexported key (`ctxRequestId`) to reduce accidental key collisions
	// with other packages.
	//
	// GetRequestId provides a controlled way for other packages to read
	// the request ID without exposing the context key itself.
	//
	// Note: This does not prevent intentional overwriting if the same key
	// is reused within our package or by code that can access that key.
	requestId := middleware.GetRequestId(ctx)

	db.logger.Info("Info", "request id", requestId)

	// pg_sleep(20) adds a 20-second delay for testing request cancellation.
	// If the client cancels the request, r.Context() signals cancellation to QueryContext,
	// allowing the database query to be cancelled instead of continuing unnecessarily.
	// Remove pg_sleep(20) after testing.
	rows, err := db.db.QueryContext(ctx, `SELECT id, title, description, price, city, created_at
			FROM listing ORDER BY created_at DESC
			LIMIT 200
			`)
	if err != nil {
		db.logger.Error("Error", "request id", requestId, "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	list := []listing{}
	for rows.Next() {
		var l listing
		err := rows.Scan(&l.Id, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt)
		if err != nil {
			db.logger.Error("Error", "request id", requestId, "err", err)
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
	ctx := r.Context()
	id := r.PathValue("id")

	_, err := db.db.ExecContext(ctx, `DELETE FROM listing where id=$1`, id)

	if err != nil {
		log.Printf("error from whene user delete some of information %v", err)
		// http.Error(w, "inter server error", http.StatusInternalServerError)

		// create custome error
		httpx.Error(w, http.StatusInternalServerError, "Something went wrong", httpx.InvalidID())
		return
	}
	w.WriteHeader(http.StatusOK)
}

// read mass assigment

func (db dbStruct) CreateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contextId := middleware.GetRequestId(ctx)
	var list CreateListingRequest
	var res SendRequestForCreateListing
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		db.logger.Error("error", "resquest id", contextId, "error", err)
		httpx.Error(w, http.StatusBadRequest, "Invlid JSON", httpx.BadRequest())
		return
	}
	if err := list.Validator(); err != nil {
		var validatorErr ValidatorError
		errors.As(err, &validatorErr)
		db.logger.Error("error", "request_id", contextId, "error", err)
		httpx.ValidatorError(w, http.StatusUnprocessableEntity, err.Error(), httpx.ValidationError(), validatorErr.Field)
		return
	}
	row := db.db.QueryRowContext(ctx, `INSERT INTO listing (title,description,price,city) 
	VALUES ($1,$2,$3,$4) 
	RETURNING id`, list.Title, list.Description, list.Price, list.City)

	scanErr := row.Scan(&res.Id)
	if scanErr != nil {
		db.logger.Error("error", "resquest id", contextId, "error", scanErr)
		httpx.Error(w, http.StatusInternalServerError, "Something went wrong", httpx.ListingNotFound())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		db.logger.Error("failed to encode response", "error", err)
	}
}
