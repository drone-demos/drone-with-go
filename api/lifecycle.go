package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"os"
)

type AppLifecycleResource struct{}

// Routes creates a REST router for the todos resource
func (self AppLifecycleResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/lifecycle", func(r chi.Router) {
		r.Post("/stop", self.stop)
	})

	return r
}

func (self AppLifecycleResource) stop(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
  os.Exit(0)
}
