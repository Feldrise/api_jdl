package gamecard

import (
	"feldrise.com/jdl/config"
	"github.com/go-chi/chi"
)

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", config.GetAll)
	router.Get("/random", config.GetRandom)
	router.Post("/", config.Create)
	router.Put("/{id}", config.Update)

	return router
}
