package database

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"uuid"

	db "github.com/AniketDubey199/PostgreSQL/db/sqlc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler struct {
	queries *db.Queries
}

func NewHandler(pool *pgxpool.Pool) *UserHandler {
	return &UserHandler{
		queries: db.New(pool),
	}
}

type CreateUserRequest struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type UpdateUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

func Migrate() {
	key := os.Getenv("DATABASE_URL")
	if key == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	m, err := migrate.New(
		"file://db/migrations", // Ensure this path matches your actual folder name
		key,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrator: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil {
		// Ignore the error if the database is already up to date
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("Database is already up to date.")
			return
		}
		log.Fatalf("Failed to apply migration: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.queries.CreateUser(r.Context(), db.CreateUserParams{
		Name:  *request.Name,
		Email: *request.Email,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	useruuid, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req UpdateUser
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.queries.UpdateUserByID(r.Context(), db.UpdateUserByIDParams{
		Email: *req.Email,
		UserID: pgtype.UUID{
			Bytes: useruuid,
			Valid: true,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode("Succesfully updated ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	useruuid, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uuiduser := pgtype.UUID{
		Bytes: useruuid,
		Valid: true,
	}
	getuser, err := h.queries.GetUserByID(r.Context(), uuiduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(&getuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	useruuid, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uuiduser := pgtype.UUID{
		Bytes: useruuid,
		Valid: true,
	}
	getuser, err := h.queries.DeleteUserById(r.Context(), uuiduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(&getuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
