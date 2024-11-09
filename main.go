package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

type Tab struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("GET /api/tabs", func(w http.ResponseWriter, r *http.Request) {
		tabs := []Tab{
			{
				"foo",
				"Foo",
			},
			{
				"bar",
				"Bar",
			},
			{
				"baz",
				"Baz",
			},
		}
		_ = json.NewEncoder(w).Encode(tabs)
	})
	http.HandleFunc("GET /api/tabs/{name}", func(w http.ResponseWriter, r *http.Request) {

	})

	slog.Info("Starting app", "url", "http://localhost:3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
