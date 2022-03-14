package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
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
		//w.Write([]byte("welcome"))
		//ещё один коммент
		w.Write(jsonResp)
	})
	http.ListenAndServe(":3000", r)

}

type TestResponse struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}
