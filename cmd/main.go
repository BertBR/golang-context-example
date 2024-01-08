package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BertBR/golang-context-example/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	h := handlers.New()

	// Api versioning
	v1Routes := r.Route("/", func(r chi.Router) {
		r.Mount("/users", h)
	})

	r.Mount("/v1", v1Routes)

	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", r)
}
