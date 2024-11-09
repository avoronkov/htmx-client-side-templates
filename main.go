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

type TabsResponse struct {
	Tabs []Tab `json:"tabs"`
}

type ContentLine struct {
	First  string `json:"first"`
	Second string `json:"second"`
	Third  string `json:"third"`
}

type TabContent struct {
	Lines []ContentLine `json:"lines"`
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/favicon.ico")
	})

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
		resp := &TabsResponse{
			Tabs: tabs,
		}
		_ = json.NewEncoder(w).Encode(resp)
	})
	http.HandleFunc("GET /api/tabs/{name}", func(w http.ResponseWriter, r *http.Request) {
		tabs := map[string]TabContent{
			"foo": {
				Lines: []ContentLine{
					{"Lorem", "ipsum", "dolor"},
					{"sit", "amet", "consectetur"},
					{"adipiscing", "elit", "sed"},
				},
			},
			"bar": {
				Lines: []ContentLine{
					{"do", "eiusmod", "tempor"},
					{"incididunt", "ut", "labore"},
					{"et", "dolore", "magna aliqua"},
				},
			},
			"baz": {
				Lines: []ContentLine{
					{"Ut enim", "ad minim", "veniam"},
					{"quis", "nostrud", "exercitation"},
					{"ullamco", "laboris", "nisi"},
				},
			},
		}
		tab := tabs[r.PathValue("name")]
		_ = json.NewEncoder(w).Encode(tab)
	})

	slog.Info("Starting app", "url", "http://localhost:3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
