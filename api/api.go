package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"time"
)

func ApiInit() {
	log.Println("Initializing API server")
	r := chi.NewRouter()
	addMiddleware(r)
  addRoutes(r)
  start(r)
}

func addMiddleware(mux *chi.Mux) {
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.URLFormat)
	mux.Use(render.SetContentType(render.ContentTypeJSON))
	mux.Use(middleware.Timeout(60 * time.Second)) // Set a timeout
}

func addRoutes(mux *chi.Mux) {
	// Add a simple resource
	mux.Mount("/api/id", AppResource{}.Routes())
	// Live a healthy life!
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yolo"))
	})
}

func start(mux *chi.Mux) {
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
