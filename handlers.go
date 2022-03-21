package main

import (
	//"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	Idtasks     uint
	Description string
}

type Comment struct {
	Idcomments uint
	Idtasks    uint
	Text       string
}

/*var db *gorm.DB

func SetDBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
		ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}*/

func service() http.Handler {
	dsn, _ := os.LookupEnv("DSN")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Task{}, &Comment{})
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	//r.Use(SetDBMiddleware)

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

		w.Write([]byte("all done.\n"))
	})

	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		res := db.Take(&Task{})
		jsonResp, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	r.Put("/tasks", func(w http.ResponseWriter, r *http.Request) {
		task := Task{Description: "bla bla"}
		db.Create(&task)
		/*jsonResp, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)*/
	})

	return r
}

type TestResponse struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}
