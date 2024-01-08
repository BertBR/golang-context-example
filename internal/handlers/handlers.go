package handlers

import (
	"fmt"
	"net/http"

	"github.com/BertBR/golang-context-example/pkg/middlewares"
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.With(middlewares.DefaultPermissions).Post("/", AddUser)
	})

	return router
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//Here you can get all user data from the context middleware and do anything you want as "Cancel/recover request"
	user := ctx.Value(struct{}{}).(middlewares.User)
	fmt.Printf("%+v\n", user)

	w.WriteHeader(http.StatusOK)
}
