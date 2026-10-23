package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	db "github.com/AniketDubey199/PostgreSQL/db/sqlc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

// Add this to your database/migrate.go file
func RollbackOne() {
	key := os.Getenv("DATABASE_URL")
	m, err := migrate.New("file://db/migrations", key)
	if err != nil {
		log.Fatalf("Failed to initialize migrator: %v", err)
	}

	// Steps down by exactly 1 version
	if err := m.Steps(-1); err != nil {
		log.Fatalf("Rollback failed: %v", err)
	}
	log.Println("Successfully rolled back 1 database version.")
}

func Create(ctx context.Context, pool *pgxpool.Pool) {
	queries := db.New(pool)
	newUser, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:  "tushar",
		Email: "aniketdubey@gmail.com",
	})
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	fmt.Printf("Success! Created user %s with email %s\n", newUser.Name, newUser.Email)
}
