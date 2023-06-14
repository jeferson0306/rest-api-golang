package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"net/http"
	"rest-api-golang/configs"
	"rest-api-golang/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
