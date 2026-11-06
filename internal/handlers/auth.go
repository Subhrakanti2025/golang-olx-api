package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Subhrakanti2025/golang-olx-api/internal/httpx"
	"github.com/Subhrakanti2025/golang-olx-api/internal/middleware"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type dbStructForAuth struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewDbConstructorForAuth(db *sql.DB, logger *slog.Logger) *dbStructForAuth {
	return &dbStructForAuth{
		db:     db,
		logger: logger,
	}
}

func (db dbStructForAuth) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := middleware.GetRequestId(ctx)

	var user UserStruct
	var userResponse UserInsertResposne

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		db.logger.Error("error", "request id", requestId, "error", err)
		httpx.Error(w, http.StatusBadRequest, "Invalid JSON", httpx.BadRequest())
	}

	validateError := user.AuthValidate()
	if validateError != nil {
		var authValidateErr *AuthValidatorErr
		errors.As(validateError, &authValidateErr)
		db.logger.Error("error", "request id", requestId, "error", validateError)
		httpx.ValidatorError(w, http.StatusBadRequest, validateError.Error(), httpx.BadRequest(), authValidateErr.Field)
		return
	}

	password := []byte(user.Password)
	encrpPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		db.logger.Error("Error", "request id", requestId, "error", err)
		httpx.Error(w, http.StatusInternalServerError, "Internal server error", httpx.InternalServerError())
		return
	}

	rows := db.db.QueryRowContext(ctx, `INSERT INTO users (name,email,password) VALUES ($1,$2,$3) RETURNING id`, user.Name, user.Email, encrpPass)

	scanErr := rows.Scan(&userResponse.Id)

	// make it correct letter have a bug
	if scanErr != nil {
		var pgErr *pgconn.PgError
		errors.As(scanErr, &pgErr)
		fmt.Println("pg error code", pgErr)
		db.logger.Error("error", "request id", requestId, "error", scanErr)
		// we will do give more specific error letter
		if pgErr.Code == "23505" {
			httpx.Error(w, http.StatusConflict, "Duplicate email entry", httpx.StatusConflict())
			return
		}
		httpx.Error(w, http.StatusInternalServerError, "Internal server error", httpx.InternalServerError())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&userResponse)
}
