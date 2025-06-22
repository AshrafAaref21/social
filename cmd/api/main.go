package main

import (
	"log"

	"github.com/ashrafaaref20/social/internal/db"
	"github.com/ashrafaaref20/social/internal/env"
	"github.com/ashrafaaref20/social/internal/store"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env (silently ignore if file is missing in production)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment")
	}
}

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpass@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Println("Database Connection pool is established.")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
