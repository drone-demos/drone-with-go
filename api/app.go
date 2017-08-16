package api

import (
	"context"
	. "drone-with-go/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type AppResource struct{}

// Routes creates a REST router for the todos resource
func (self AppResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(r chi.Router) {
		r.Use(AppCtx) // do some preprocessing of the input value
		r.Get("/", self.remember)
		r.Post("/", self.memorize)
	})

	return r
}

func (self AppResource) remember(w http.ResponseWriter, r *http.Request) {
	record := r.Context().Value("record").(*Record)
	render.Status(r, http.StatusOK)
	render.Render(w, r, NewResponse("GET", record.Value))
}

func (self AppResource) memorize(w http.ResponseWriter, r *http.Request) {
	record := r.Context().Value("record").(*Record)
	render.Status(r, http.StatusOK)
	render.Render(w, r, NewResponse("POST", record.Value))
}

// AppCtx middleware is used to load an object from
// the URL parameters passed through as the request. In case
// the object could not be found, we stop here and return a 404.
func AppCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var record *Record

		if id := chi.URLParam(r, "id"); id != "" {
			i64, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				render.Render(w, r, ErrNotANumber)
				return
			}
			// Here is where we could lookup the record...
			record = &Record{Value: i64}

		} else {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "record", record)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//--------------------------------------------------------------------------
// response renderers
//--------------------------------------------------------------------------
type Response struct {
	Operation  string `json:"op"`
	Identifier uint64 `json:"id"`
}

func (u *Response) Bind(r *http.Request) error {
	return nil
}

func (u *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewResponse(action string, id uint64) *Response {
	resp := &Response{Operation: action, Identifier: id}
	return resp
}
