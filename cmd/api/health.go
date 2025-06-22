package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Okay..!"))

	// app.store.Posts.Create(r.Context())
	log.Println("created")
}
