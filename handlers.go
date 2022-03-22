package main

import (
	//"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	//"time"

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

func service() http.Handler {
	dsn, _ := os.LookupEnv("DSN")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Task{}, &Comment{})
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
	r.Get("/tasks/{taskid}", func(w http.ResponseWriter, r *http.Request) {
		taskid := chi.URLParam(r, "taskid")
		task := Task{}
		db.First(&task, taskid)
		jsonResp, err := json.Marshal(task)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {

		tasks := []Task{}
		db.Find(&tasks)
		jsonResp, err := json.Marshal(tasks)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
		task := Task{}
		json.NewDecoder(r.Body).Decode(&task)
		db.Create(&task)
	})

	return r
}

type TestResponse struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}
