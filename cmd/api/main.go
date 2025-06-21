package main

import (
	"log"

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
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))
}