package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func service() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var reply_struct TestResponse
		reply_struct.Field1 = "abc"
		reply_struct.Field2 = "def"
		jsonResp, err := json.Marshal(reply_struct)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
		// Simulates some hard work.
		//
		// We want this handler to complete successfully during a shutdown signal,
		// so consider the work here as some background routine to fetch a long running
		// search query to find as many results as possible, but, instead we cut it short
		// and respond with what we have so far. How a shutdown is handled is entirely
		// up to the developer, as some code blocks are preemptable, and others are not.
		time.Sleep(5 * time.Second)

		w.Write([]byte(fmt.Sprintf("all done.\n")))
	})

	return r
}

type TestResponse struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}
