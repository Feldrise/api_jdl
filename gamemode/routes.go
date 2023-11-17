package gamemode

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
	router.Post("/", config.Create)

	return router
}
