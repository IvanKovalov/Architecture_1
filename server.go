package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Date struct {
	Hour time.Time `json:"Поточний час"`
}

func main() {
	r := chi.NewRouter()
	r.Get("/time", func(w http.ResponseWriter, r *http.Request) {
		data := Date{time.Now()}
		js, _ := json.Marshal(data)
		w.Write(js)
	})
	http.ListenAndServe(":8795", r)
}

//Test 1
