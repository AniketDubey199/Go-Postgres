package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AniketDubey199/PostgreSQL/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// rollback := flag.Bool("rollback", false, "Rollback data by step 1")  // this is for the rollback using the flag tool , we will use the make file version too
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	// if *rollback {
	// 	database.RollbackOne()
	// 	return
	// }
	database.Migrate()

	ctx := context.Background()
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}

	config.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		fmt.Println("New connnection established")
		return err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create connection pool %v", err)
	}

	database.Create(ctx, pool)

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to create connection pool %v", err)
	}
	fmt.Println("Database connection pool intialized successfully ")
}
